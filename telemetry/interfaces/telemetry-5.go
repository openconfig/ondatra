package interfaces

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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Unnumbered", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_RoutedVlan_Ipv6_Unnumbered
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Unnumbered) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_VlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_VlanPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_VlanPath) Get(t testing.TB) oc.Interface_RoutedVlan_Vlan_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_VlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Vlan_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_VlanPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a ONCE subscription.
func (n *Interface_RoutedVlan_VlanPathAny) Get(t testing.TB) []oc.Interface_RoutedVlan_Vlan_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Interface_RoutedVlan_Vlan_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_VlanPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Vlan_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_VlanPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool) *oc.Interface_RoutedVlan_Vlan_UnionWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Vlan_UnionWatcher{}
	gs := &oc.Interface_RoutedVlan{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_VlanPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Vlan_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_VlanPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool) *oc.Interface_RoutedVlan_Vlan_UnionWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_VlanPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_VlanPath) Await(t testing.TB, timeout time.Duration, val oc.Interface_RoutedVlan_Vlan_Union) *oc.QualifiedInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan to the batch object.
func (n *Interface_RoutedVlan_VlanPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_VlanPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Vlan_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_VlanPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool) *oc.Interface_RoutedVlan_Vlan_UnionWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Vlan_UnionWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_VlanPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Vlan_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_VlanPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Vlan_Union) bool) *oc.Interface_RoutedVlan_Vlan_UnionWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_VlanPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/state/vlan to the batch object.
func (n *Interface_RoutedVlan_VlanPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_VlanPath extracts the value of the leaf Vlan from its parent oc.Interface_RoutedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedInterface_RoutedVlan_Vlan_Union.
func convertInterface_RoutedVlan_VlanPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan) *oc.QualifiedInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	qv := &oc.QualifiedInterface_RoutedVlan_Vlan_Union{
		Metadata: md,
	}
	val := parent.Vlan
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_SubinterfacePath) Get(t testing.TB) *oc.Interface_Subinterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a ONCE subscription.
func (n *Interface_SubinterfacePathAny) Get(t testing.TB) []*oc.Interface_Subinterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface)))
		return false
	})
	return c
}

func watch_Interface_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface) bool) *oc.Interface_SubinterfaceWatcher {
	t.Helper()
	w := &oc.Interface_SubinterfaceWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface) bool) *oc.Interface_SubinterfaceWatcher {
	t.Helper()
	return watch_Interface_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface) *oc.QualifiedInterface_Subinterface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface to the batch object.
func (n *Interface_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_SubinterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface) bool) *oc.Interface_SubinterfaceWatcher {
	t.Helper()
	w := &oc.Interface_SubinterfaceWatcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface) bool) *oc.Interface_SubinterfaceWatcher {
	t.Helper()
	return watch_Interface_SubinterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface to the batch object.
func (n *Interface_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_AdminStatusPath) Lookup(t testing.TB) *oc.QualifiedE_Interface_AdminStatus {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_AdminStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_AdminStatusPath) Get(t testing.TB) oc.E_Interface_AdminStatus {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_AdminStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Interface_AdminStatus {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Interface_AdminStatus
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_AdminStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a ONCE subscription.
func (n *Interface_Subinterface_AdminStatusPathAny) Get(t testing.TB) []oc.E_Interface_AdminStatus {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Interface_AdminStatus
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_AdminStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Interface_AdminStatus {
	t.Helper()
	c := &oc.CollectionE_Interface_AdminStatus{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Interface_AdminStatus) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_AdminStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Interface_AdminStatus) bool) *oc.E_Interface_AdminStatusWatcher {
	t.Helper()
	w := &oc.E_Interface_AdminStatusWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_AdminStatusPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Interface_AdminStatus)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_AdminStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Interface_AdminStatus) bool) *oc.E_Interface_AdminStatusWatcher {
	t.Helper()
	return watch_Interface_Subinterface_AdminStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_AdminStatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_Interface_AdminStatus) *oc.QualifiedE_Interface_AdminStatus {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Interface_AdminStatus) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status to the batch object.
func (n *Interface_Subinterface_AdminStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_AdminStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Interface_AdminStatus {
	t.Helper()
	c := &oc.CollectionE_Interface_AdminStatus{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Interface_AdminStatus) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_AdminStatusPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Interface_AdminStatus) bool) *oc.E_Interface_AdminStatusWatcher {
	t.Helper()
	w := &oc.E_Interface_AdminStatusWatcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_AdminStatusPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Interface_AdminStatus)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_AdminStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Interface_AdminStatus) bool) *oc.E_Interface_AdminStatusWatcher {
	t.Helper()
	return watch_Interface_Subinterface_AdminStatusPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/admin-status to the batch object.
func (n *Interface_Subinterface_AdminStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_AdminStatusPath extracts the value of the leaf AdminStatus from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Interface_AdminStatus.
func convertInterface_Subinterface_AdminStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedE_Interface_AdminStatus {
	t.Helper()
	qv := &oc.QualifiedE_Interface_AdminStatus{
		Metadata: md,
	}
	val := parent.AdminStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_CountersPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Counters {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_CountersPath) Get(t testing.TB) *oc.Interface_Subinterface_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a ONCE subscription.
func (n *Interface_Subinterface_CountersPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Counters {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Counters)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Counters) bool) *oc.Interface_Subinterface_CountersWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_CountersWatcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Counters) bool) *oc.Interface_Subinterface_CountersWatcher {
	t.Helper()
	return watch_Interface_Subinterface_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Counters) *oc.QualifiedInterface_Subinterface_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters to the batch object.
func (n *Interface_Subinterface_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Counters {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Counters) bool) *oc.Interface_Subinterface_CountersWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_CountersWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Counters) bool) *oc.Interface_Subinterface_CountersWatcher {
	t.Helper()
	return watch_Interface_Subinterface_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters to the batch object.
func (n *Interface_Subinterface_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_CarrierTransitionsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_CarrierTransitionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a ONCE subscription.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_CarrierTransitionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_CarrierTransitionsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_CarrierTransitionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions to the batch object.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_CarrierTransitionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_CarrierTransitionsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_CarrierTransitionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/carrier-transitions to the batch object.
func (n *Interface_Subinterface_Counters_CarrierTransitionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_CarrierTransitionsPath extracts the value of the leaf CarrierTransitions from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_CarrierTransitionsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CarrierTransitions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InBroadcastPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InBroadcastPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InBroadcastPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InBroadcastPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InBroadcastPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InBroadcastPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InBroadcastPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InBroadcastPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-broadcast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InBroadcastPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InBroadcastPktsPath extracts the value of the leaf InBroadcastPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InBroadcastPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InBroadcastPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InDiscardsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InDiscardsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InDiscardsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InDiscardsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InDiscardsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InDiscardsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InDiscardsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InDiscardsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InDiscardsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InDiscardsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InDiscardsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InDiscardsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards to the batch object.
func (n *Interface_Subinterface_Counters_InDiscardsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InDiscardsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InDiscardsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InDiscardsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InDiscardsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InDiscardsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-discards to the batch object.
func (n *Interface_Subinterface_Counters_InDiscardsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InDiscardsPath extracts the value of the leaf InDiscards from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InDiscardsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InDiscards
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors to the batch object.
func (n *Interface_Subinterface_Counters_InErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-errors to the batch object.
func (n *Interface_Subinterface_Counters_InErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InErrorsPath extracts the value of the leaf InErrors from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InFcsErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InFcsErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InFcsErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InFcsErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InFcsErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InFcsErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InFcsErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InFcsErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InFcsErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InFcsErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InFcsErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InFcsErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors to the batch object.
func (n *Interface_Subinterface_Counters_InFcsErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InFcsErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InFcsErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InFcsErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InFcsErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InFcsErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-fcs-errors to the batch object.
func (n *Interface_Subinterface_Counters_InFcsErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InFcsErrorsPath extracts the value of the leaf InFcsErrors from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InFcsErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFcsErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InMulticastPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InMulticastPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InMulticastPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InMulticastPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InMulticastPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InMulticastPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InMulticastPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InMulticastPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InMulticastPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InMulticastPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InMulticastPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InMulticastPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InMulticastPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InMulticastPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InMulticastPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InMulticastPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InMulticastPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InMulticastPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-multicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InMulticastPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InMulticastPktsPath extracts the value of the leaf InMulticastPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InMulticastPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InMulticastPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets to the batch object.
func (n *Interface_Subinterface_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-octets to the batch object.
func (n *Interface_Subinterface_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InUnicastPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InUnicastPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InUnicastPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InUnicastPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InUnicastPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InUnicastPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InUnicastPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InUnicastPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InUnicastPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InUnicastPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InUnicastPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InUnicastPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InUnicastPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InUnicastPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InUnicastPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InUnicastPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InUnicastPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InUnicastPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_InUnicastPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InUnicastPktsPath extracts the value of the leaf InUnicastPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InUnicastPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InUnicastPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_InUnknownProtosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_InUnknownProtosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_InUnknownProtosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_InUnknownProtosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_InUnknownProtosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a ONCE subscription.
func (n *Interface_Subinterface_Counters_InUnknownProtosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InUnknownProtosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InUnknownProtosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_InUnknownProtosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InUnknownProtosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InUnknownProtosPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_InUnknownProtosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos to the batch object.
func (n *Interface_Subinterface_Counters_InUnknownProtosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_InUnknownProtosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_InUnknownProtosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_InUnknownProtosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_InUnknownProtosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_InUnknownProtosPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/in-unknown-protos to the batch object.
func (n *Interface_Subinterface_Counters_InUnknownProtosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_InUnknownProtosPath extracts the value of the leaf InUnknownProtos from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_InUnknownProtosPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InUnknownProtos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_LastClearPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_LastClearPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_LastClearPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_LastClearPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_LastClearPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a ONCE subscription.
func (n *Interface_Subinterface_Counters_LastClearPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_LastClearPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_LastClearPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_LastClearPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_LastClearPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_LastClearPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_LastClearPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear to the batch object.
func (n *Interface_Subinterface_Counters_LastClearPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_LastClearPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_LastClearPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_LastClearPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_LastClearPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_LastClearPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/last-clear to the batch object.
func (n *Interface_Subinterface_Counters_LastClearPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_LastClearPath extracts the value of the leaf LastClear from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_LastClearPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastClear
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutBroadcastPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutBroadcastPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutBroadcastPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutBroadcastPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutBroadcastPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutBroadcastPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutBroadcastPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutBroadcastPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-broadcast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutBroadcastPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutBroadcastPktsPath extracts the value of the leaf OutBroadcastPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutBroadcastPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutBroadcastPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutDiscardsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutDiscardsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutDiscardsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutDiscardsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutDiscardsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutDiscardsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutDiscardsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutDiscardsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutDiscardsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutDiscardsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutDiscardsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutDiscardsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards to the batch object.
func (n *Interface_Subinterface_Counters_OutDiscardsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutDiscardsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutDiscardsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutDiscardsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutDiscardsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutDiscardsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-discards to the batch object.
func (n *Interface_Subinterface_Counters_OutDiscardsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutDiscardsPath extracts the value of the leaf OutDiscards from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutDiscardsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutDiscards
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors to the batch object.
func (n *Interface_Subinterface_Counters_OutErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-errors to the batch object.
func (n *Interface_Subinterface_Counters_OutErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutErrorsPath extracts the value of the leaf OutErrors from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutMulticastPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutMulticastPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutMulticastPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutMulticastPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutMulticastPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutMulticastPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutMulticastPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutMulticastPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-multicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutMulticastPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutMulticastPktsPath extracts the value of the leaf OutMulticastPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutMulticastPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMulticastPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets to the batch object.
func (n *Interface_Subinterface_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-octets to the batch object.
func (n *Interface_Subinterface_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Counters_OutUnicastPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Counters_OutUnicastPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutUnicastPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Counters_OutUnicastPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutUnicastPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Counters_OutUnicastPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Counters_OutUnicastPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Counters_OutUnicastPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/counters/out-unicast-pkts to the batch object.
func (n *Interface_Subinterface_Counters_OutUnicastPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Counters_OutUnicastPktsPath extracts the value of the leaf OutUnicastPkts from its parent oc.Interface_Subinterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Counters_OutUnicastPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutUnicastPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_CpuPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_CpuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_CpuPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_CpuPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_CpuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a ONCE subscription.
func (n *Interface_Subinterface_CpuPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_CpuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_CpuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_CpuPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_CpuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_CpuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_CpuPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu to the batch object.
func (n *Interface_Subinterface_CpuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_CpuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_CpuPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_CpuPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_CpuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_CpuPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/cpu to the batch object.
func (n *Interface_Subinterface_CpuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_CpuPath extracts the value of the leaf Cpu from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_CpuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Cpu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a ONCE subscription.
func (n *Interface_Subinterface_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_DescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description to the batch object.
func (n *Interface_Subinterface_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_DescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_DescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_DescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/description to the batch object.
func (n *Interface_Subinterface_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_DescriptionPath extracts the value of the leaf Description from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a ONCE subscription.
func (n *Interface_Subinterface_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled to the batch object.
func (n *Interface_Subinterface_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_EnabledPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/enabled to the batch object.
func (n *Interface_Subinterface_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_IfindexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_IfindexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_IfindexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_IfindexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_IfindexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a ONCE subscription.
func (n *Interface_Subinterface_IfindexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_IfindexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_IfindexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_IfindexPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_IfindexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_IfindexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_IfindexPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex to the batch object.
func (n *Interface_Subinterface_IfindexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_IfindexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_IfindexPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_IfindexPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_IfindexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_IfindexPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/ifindex to the batch object.
func (n *Interface_Subinterface_IfindexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_IfindexPath extracts the value of the leaf Ifindex from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_IfindexPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Ifindex
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_IndexPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetIndex())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_IndexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a ONCE subscription.
func (n *Interface_Subinterface_IndexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_IndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_IndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_IndexPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_IndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_IndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_IndexPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index to the batch object.
func (n *Interface_Subinterface_IndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_IndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_IndexPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_IndexPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_IndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_IndexPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/index to the batch object.
func (n *Interface_Subinterface_IndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_IndexPath extracts the value of the leaf Index from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv4 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4Path) Get(t testing.TB) *oc.Interface_Subinterface_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4PathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv4 {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv4) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv4{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv4)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4) bool) *oc.Interface_Subinterface_Ipv4Watcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv4Watcher{}
	gs := &oc.Interface_Subinterface_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv4{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4) bool) *oc.Interface_Subinterface_Ipv4Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4Path) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv4) *oc.QualifiedInterface_Subinterface_Ipv4 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv4) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 to the batch object.
func (n *Interface_Subinterface_Ipv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv4 {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv4) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4) bool) *oc.Interface_Subinterface_Ipv4Watcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv4Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv4{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4) bool) *oc.Interface_Subinterface_Ipv4Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4 to the batch object.
func (n *Interface_Subinterface_Ipv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv4_Address {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Address", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv4_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_AddressPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv4_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv4_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv4_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv4_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_AddressPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv4_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv4_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv4_Address {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv4_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv4_Address{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv4_Address)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool) *oc.Interface_Subinterface_Ipv4_AddressWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv4_AddressWatcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv4_Address{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv4_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool) *oc.Interface_Subinterface_Ipv4_AddressWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_AddressPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv4_Address) *oc.QualifiedInterface_Subinterface_Ipv4_Address {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address to the batch object.
func (n *Interface_Subinterface_Ipv4_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv4_Address {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv4_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool) *oc.Interface_Subinterface_Ipv4_AddressWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv4_AddressWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv4_Address{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv4_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Address) bool) *oc.Interface_Subinterface_Ipv4_AddressWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address to the batch object.
func (n *Interface_Subinterface_Ipv4_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv4_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv4_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Address_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Address_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv4_Address_IpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Address_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Address_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_Address_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip to the batch object.
func (n *Interface_Subinterface_Ipv4_Address_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Address_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Address_IpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv4_Address_IpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Address_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Address_IpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/ip to the batch object.
func (n *Interface_Subinterface_Ipv4_Address_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv4_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_Subinterface_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv4_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv4_Address) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Address_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv4_Address_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Address_OriginPath) Get(t testing.TB) oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Address_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_IpAddressOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv4_Address_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Address_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_IpAddressOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Address_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Address_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv4_Address_OriginPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Address_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Address_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_Address_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_IpAddressOrigin) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin to the batch object.
func (n *Interface_Subinterface_Ipv4_Address_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Address_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Address_OriginPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv4_Address_OriginPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Address_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Address_OriginPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/origin to the batch object.
func (n *Interface_Subinterface_Ipv4_Address_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv4_Address_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_Subinterface_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_IpAddressOrigin.
func convertInterface_Subinterface_Ipv4_Address_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv4_Address) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_IpAddressOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv4_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv4_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Address_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv4_Address_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Address_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length to the batch object.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv4_Address_PrefixLengthPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/addresses/address/state/prefix-length to the batch object.
func (n *Interface_Subinterface_Ipv4_Address_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv4_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_Subinterface_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv4_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv4_Address) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
