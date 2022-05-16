package interfaces

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"

	config "github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriorityDecrement())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath extracts the value of the leaf PriorityDecrement from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PriorityDecrement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath extracts the value of the leaf TrackInterface from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedStringSlice {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPreemptDelay())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath extracts the value of the leaf PreemptDelay from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPreempt())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath extracts the value of the leaf Preempt from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriority())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath extracts the value of the leaf Priority from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath extracts the value of the leaf VirtualAddress from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedStringSlice {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath extracts the value of the leaf VirtualRouterId from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDhcpClient())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/dhcp-client in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_DhcpClientPath extracts the value of the leaf DhcpClient from its parent oc.Interface_RoutedVlan_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_RoutedVlan_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/config/mtu in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface_RoutedVlan_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv4_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Mtu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Neighbor) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Neighbor) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Neighbor) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Neighbor) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv4_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/config/link-layer-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath extracts the value of the leaf LinkLayerAddress from its parent oc.Interface_RoutedVlan_Ipv4_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_ProxyArp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_ProxyArp) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_ProxyArp) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_ProxyArp) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_ProxyArp) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Lookup(t testing.TB) *oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_ProxyArp_Mode{
		Metadata: md,
	}).SetVal(goStruct.GetMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Get(t testing.TB) oc.E_ProxyArp_Mode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_ProxyArp_Mode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Get(t testing.TB) []oc.E_ProxyArp_Mode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_ProxyArp_Mode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Replace(t testing.TB, val oc.E_ProxyArp_Mode) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_ProxyArp_Mode) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Update(t testing.TB, val oc.E_ProxyArp_Mode) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/config/mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_ProxyArp_Mode) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath extracts the value of the leaf Mode from its parent oc.Interface_RoutedVlan_Ipv4_ProxyArp
// and combines the update with an existing Metadata to return a *oc.QualifiedE_ProxyArp_Mode.
func convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_ProxyArp) *oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	qv := &oc.QualifiedE_ProxyArp_Mode{
		Metadata: md,
	}
	val := parent.Mode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Unnumbered
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_RoutedVlan_Ipv4_Unnumbered
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Unnumbered) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6Path) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6PathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6.
func (n *Interface_RoutedVlan_Ipv6Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 in the given batch object.
func (n *Interface_RoutedVlan_Ipv6Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6.
func (n *Interface_RoutedVlan_Ipv6Path) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 in the given batch object.
func (n *Interface_RoutedVlan_Ipv6Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6.
func (n *Interface_RoutedVlan_Ipv6Path) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 in the given batch object.
func (n *Interface_RoutedVlan_Ipv6Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Address) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Address) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Address) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Address) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv6_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv6_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_RoutedVlan_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAcceptMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath extracts the value of the leaf AcceptMode from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AcceptMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAdvertisementInterval())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath extracts the value of the leaf AdvertisementInterval from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AdvertisementInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriorityDecrement())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath extracts the value of the leaf PriorityDecrement from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PriorityDecrement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPreemptDelay())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptDelayPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPreempt())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PreemptPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriority())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_PriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_DhcpClientPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDhcpClient())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_DhcpClientPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dhcp-client in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_DhcpClientPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetDupAddrDetectTransmits())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/dup-addr-detect-transmits in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_DupAddrDetectTransmitsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/enabled in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_MtuPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/config/mtu in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_MtuPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Neighbor", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Neighbor) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Neighbor) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv6_Neighbor) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/neighbors/neighbor in the given batch object.
func (n *Interface_RoutedVlan_Ipv6_NeighborPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv6_Neighbor) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
