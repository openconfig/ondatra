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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath extracts the value of the leaf TrackInterface from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.TrackInterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPreemptDelay())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath extracts the value of the leaf PreemptDelay from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PreemptDelay
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPreempt())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/preempt to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath extracts the value of the leaf Preempt from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Preempt
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriority())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/priority to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath extracts the value of the leaf Priority from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath extracts the value of the leaf VirtualAddress from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.VirtualAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath extracts the value of the leaf VirtualLinkLocal from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.VirtualLinkLocal
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath extracts the value of the leaf VirtualRouterId from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.VirtualRouterId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_CountersPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_CountersPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_CountersPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Counters)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters) bool) *oc.Interface_RoutedVlan_Ipv6_CountersWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_CountersWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters) bool) *oc.Interface_RoutedVlan_Ipv6_CountersWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters to the batch object.
func (n *Interface_RoutedVlan_Ipv6_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Counters) bool) *oc.Interface_RoutedVlan_Ipv6_CountersWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_CountersPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters to the batch object.
func (n *Interface_RoutedVlan_Ipv6_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath extracts the value of the leaf InDiscardedPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_InDiscardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InDiscardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-error-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InErrorPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_InErrorPktsPath extracts the value of the leaf InErrorPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_InErrorPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InErrorPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath extracts the value of the leaf InForwardedOctets from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_InForwardedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InForwardedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InForwardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath extracts the value of the leaf InForwardedPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_InForwardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InForwardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_InOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_InPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_InPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/in-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath extracts the value of the leaf OutDiscardedPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_OutDiscardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutDiscardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-error-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutErrorPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath extracts the value of the leaf OutErrorPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_OutErrorPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutErrorPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath extracts the value of the leaf OutForwardedOctets from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutForwardedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath extracts the value of the leaf OutForwardedPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_OutForwardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutForwardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Counters_OutPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/counters/out-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Interface_RoutedVlan_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv6_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_DhcpClientPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDhcpClient())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_DhcpClientPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_DhcpClientPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_DhcpClientPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_DhcpClientPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client to the batch object.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_DhcpClientPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dhcp-client to the batch object.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_DhcpClientPath extracts the value of the leaf DhcpClient from its parent oc.Interface_RoutedVlan_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_DhcpClientPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.DhcpClient
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetDupAddrDetectTransmits())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits to the batch object.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/dup-addr-detect-transmits to the batch object.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath extracts the value of the leaf DupAddrDetectTransmits from its parent oc.Interface_RoutedVlan_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.DupAddrDetectTransmits
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_RoutedVlan_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_MtuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_MtuPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_MtuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu to the batch object.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_MtuPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/state/mtu to the batch object.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface_RoutedVlan_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv6_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Mtu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_NeighborPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Neighbor)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor) bool) *oc.Interface_RoutedVlan_Ipv6_NeighborWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_NeighborWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor) bool) *oc.Interface_RoutedVlan_Ipv6_NeighborWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Neighbor) *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor to the batch object.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor) bool) *oc.Interface_RoutedVlan_Ipv6_NeighborWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_NeighborPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor to the batch object.
func (n *Interface_RoutedVlan_Ipv6_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Neighbor_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Neighbor_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Neighbor_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Neighbor_IpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_IpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Neighbor_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv6_Neighbor_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/is-router to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_IsRouterPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Neighbor_IsRouterPath extracts the value of the leaf IsRouter from its parent oc.Interface_RoutedVlan_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_Neighbor_IsRouterPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Neighbor) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.IsRouter
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/link-layer-address to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath extracts the value of the leaf LinkLayerAddress from its parent oc.Interface_RoutedVlan_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv6_Neighbor_LinkLayerAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LinkLayerAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath) Lookup(t testing.TB) *oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath) Get(t testing.TB) oc.E_Neighbor_NeighborState {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Neighbor_NeighborState
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePathAny) Get(t testing.TB) []oc.E_Neighbor_NeighborState {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Neighbor_NeighborState
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Neighbor_NeighborState {
	t.Helper()
	c := &oc.CollectionE_Neighbor_NeighborState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Neighbor_NeighborState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Neighbor_NeighborState) bool) *oc.E_Neighbor_NeighborStateWatcher {
	t.Helper()
	w := &oc.E_Neighbor_NeighborStateWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Neighbor_NeighborState)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Neighbor_NeighborState) bool) *oc.E_Neighbor_NeighborStateWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath) Await(t testing.TB, timeout time.Duration, val oc.E_Neighbor_NeighborState) *oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Neighbor_NeighborState) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Neighbor_NeighborState {
	t.Helper()
	c := &oc.CollectionE_Neighbor_NeighborState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Neighbor_NeighborState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Neighbor_NeighborState) bool) *oc.E_Neighbor_NeighborStateWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/neighbor-state to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_NeighborStatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath extracts the value of the leaf NeighborState from its parent oc.Interface_RoutedVlan_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Neighbor_NeighborState.
func convertInterface_RoutedVlan_Ipv6_Neighbor_NeighborStatePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Neighbor) *oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	qv := &oc.QualifiedE_Neighbor_NeighborState{
		Metadata: md,
	}
	val := parent.NeighborState
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Neighbor_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPath) Get(t testing.TB) oc.E_IfIp_NeighborOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_NeighborOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Neighbor_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_NeighborOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_NeighborOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_NeighborOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_NeighborOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Neighbor_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_NeighborOriginWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Neighbor_OriginPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_NeighborOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_NeighborOrigin) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_NeighborOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_NeighborOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Neighbor_OriginPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Neighbor_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Neighbor_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_RoutedVlan_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_NeighborOrigin.
func convertInterface_RoutedVlan_Ipv6_Neighbor_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Neighbor) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_NeighborOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_RouterAdvertisementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) bool) *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) bool) *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement) *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) bool) *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisementPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/interval to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath extracts the value of the leaf Interval from its parent oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/lifetime to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath extracts the value of the leaf Lifetime from its parent oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_LifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Lifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSuppress())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_RouterAdvertisement", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/router-advertisement/state/suppress to the batch object.
func (n *Interface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath extracts the value of the leaf Suppress from its parent oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_RouterAdvertisement_SuppressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_RouterAdvertisement) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Suppress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Unnumbered", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Unnumbered
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Unnumbered)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_UnnumberedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) bool) *oc.Interface_RoutedVlan_Ipv6_UnnumberedWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_UnnumberedWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) bool) *oc.Interface_RoutedVlan_Ipv6_UnnumberedWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_UnnumberedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Unnumbered) *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered to the batch object.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Unnumbered{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) bool) *oc.Interface_RoutedVlan_Ipv6_UnnumberedWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_UnnumberedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered to the batch object.
func (n *Interface_RoutedVlan_Ipv6_UnnumberedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Unnumbered", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Unnumbered_EnabledPath(t, n, timeout, predicate)
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
