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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Replace(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Update(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_EgressMapping_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Get(t testing.TB) oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_VlanStackAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Get(t testing.TB) []oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_VlanStackAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Replace(t testing.TB, val oc.E_VlanTypes_VlanStackAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanStackAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Update(t testing.TB, val oc.E_VlanTypes_VlanStackAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-stack-action in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanStackAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath extracts the value of the leaf VlanStackAction from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_VlanStackAction.
func convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_VlanStackAction{
		Metadata: md,
	}
	val := parent.VlanStackAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_IngressMapping{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_IngressMapping
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_IngressMapping{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_IngressMapping
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_IngressMapping) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_IngressMapping) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_IngressMapping) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_IngressMapping) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Replace(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Update(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_IngressMapping_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface_Subinterface_Vlan_IngressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_IngressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Get(t testing.TB) oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_VlanStackAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Get(t testing.TB) []oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_VlanStackAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Replace(t testing.TB, val oc.E_VlanTypes_VlanStackAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanStackAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Update(t testing.TB, val oc.E_VlanTypes_VlanStackAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/config/vlan-stack-action in the given batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanStackAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath extracts the value of the leaf VlanStackAction from its parent oc.Interface_Subinterface_Vlan_IngressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_VlanStackAction.
func convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_VlanStackAction{
		Metadata: md,
	}
	val := parent.VlanStackAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_MatchPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_MatchPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match.
func (n *Interface_Subinterface_Vlan_MatchPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match in the given batch object.
func (n *Interface_Subinterface_Vlan_MatchPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match.
func (n *Interface_Subinterface_Vlan_MatchPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match in the given batch object.
func (n *Interface_Subinterface_Vlan_MatchPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match.
func (n *Interface_Subinterface_Vlan_MatchPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match in the given batch object.
func (n *Interface_Subinterface_Vlan_MatchPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Replace(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Update(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/inner-vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath extracts the value of the leaf InnerVlanIds from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.InnerVlanIds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath extracts the value of the leaf OuterVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath extracts the value of the leaf InnerHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/inner-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath extracts the value of the leaf InnerLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath extracts the value of the leaf OuterHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/config/outer-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath extracts the value of the leaf OuterLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath extracts the value of the leaf InnerHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/inner-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath extracts the value of the leaf InnerLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Replace(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Update(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath extracts the value of the leaf OuterVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.OuterVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath extracts the value of the leaf InnerVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Replace(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Update(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/config/outer-vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath extracts the value of the leaf OuterVlanIds from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.OuterVlanIds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath extracts the value of the leaf InnerVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath extracts the value of the leaf OuterHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/config/outer-low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath extracts the value of the leaf OuterLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTagged
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/inner-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath extracts the value of the leaf InnerVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTagged
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/config/outer-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath extracts the value of the leaf OuterVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTagged
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedList
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Replace(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Update(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/config/vlan-ids in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath extracts the value of the leaf VlanIds from its parent oc.Interface_Subinterface_Vlan_Match_SingleTaggedList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.VlanIds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_SingleTagged
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_SingleTagged) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_SingleTagged) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_SingleTagged) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_SingleTagged) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/high-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath extracts the value of the leaf HighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.HighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/config/low-vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath extracts the value of the leaf LowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.LowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_Match_SingleTagged
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTagged) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Get(t testing.TB) oc.Interface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Get(t testing.TB) []oc.Interface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Interface_Subinterface_Vlan_VlanId_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_VlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Replace(t testing.TB, val oc.Interface_Subinterface_Vlan_VlanId_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_VlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Interface_Subinterface_Vlan_VlanId_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Update(t testing.TB, val oc.Interface_Subinterface_Vlan_VlanId_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_VlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Interface_Subinterface_Vlan_VlanId_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan
// and combines the update with an existing Metadata to return a *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union.
func convertInterface_Subinterface_Vlan_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan) *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	qv := &oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_TpidPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}).SetVal(goStruct.GetTpid())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/tpid with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/tpid with a ONCE subscription.
func (n *Interface_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/tpid.
func (n *Interface_TpidPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/tpid in the given batch object.
func (n *Interface_TpidPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/tpid.
func (n *Interface_TpidPath) Replace(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/tpid in the given batch object.
func (n *Interface_TpidPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/tpid.
func (n *Interface_TpidPath) Update(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/tpid in the given batch object.
func (n *Interface_TpidPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_TypePath) Lookup(t testing.TB) *oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_TypePath) Get(t testing.TB) oc.E_IETFInterfaces_InterfaceType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IETFInterfaces_InterfaceType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/type with a ONCE subscription.
func (n *Interface_TypePathAny) Get(t testing.TB) []oc.E_IETFInterfaces_InterfaceType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IETFInterfaces_InterfaceType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/type.
func (n *Interface_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/type in the given batch object.
func (n *Interface_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/type.
func (n *Interface_TypePath) Replace(t testing.TB, val oc.E_IETFInterfaces_InterfaceType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/type in the given batch object.
func (n *Interface_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_IETFInterfaces_InterfaceType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/type.
func (n *Interface_TypePath) Update(t testing.TB, val oc.E_IETFInterfaces_InterfaceType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/type in the given batch object.
func (n *Interface_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_IETFInterfaces_InterfaceType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_TypePath extracts the value of the leaf Type from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IETFInterfaces_InterfaceType.
func convertInterface_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	qv := &oc.QualifiedE_IETFInterfaces_InterfaceType{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
