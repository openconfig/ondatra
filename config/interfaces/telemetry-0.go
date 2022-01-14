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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *InterfacePath) Lookup(t testing.TB) *oc.QualifiedInterface {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *InterfacePath) Get(t testing.TB) *oc.Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface with a ONCE subscription.
func (n *InterfacePathAny) Get(t testing.TB) []*oc.Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface.
func (n *InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface in the given batch object.
func (n *InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface.
func (n *InterfacePath) Replace(t testing.TB, val *oc.Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface in the given batch object.
func (n *InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface.
func (n *InterfacePath) Update(t testing.TB, val *oc.Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface in the given batch object.
func (n *InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_AggregationPath) Lookup(t testing.TB) *oc.QualifiedInterface_Aggregation {
	t.Helper()
	goStruct := &oc.Interface_Aggregation{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Aggregation{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_AggregationPath) Get(t testing.TB) *oc.Interface_Aggregation {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_AggregationPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Aggregation {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Aggregation
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Aggregation{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation with a ONCE subscription.
func (n *Interface_AggregationPathAny) Get(t testing.TB) []*oc.Interface_Aggregation {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Aggregation
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation.
func (n *Interface_AggregationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation in the given batch object.
func (n *Interface_AggregationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation.
func (n *Interface_AggregationPath) Replace(t testing.TB, val *oc.Interface_Aggregation) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation in the given batch object.
func (n *Interface_AggregationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Aggregation) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation.
func (n *Interface_AggregationPath) Update(t testing.TB, val *oc.Interface_Aggregation) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation in the given batch object.
func (n *Interface_AggregationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Aggregation) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_LagTypePath) Lookup(t testing.TB) *oc.QualifiedE_IfAggregate_AggregationType {
	t.Helper()
	goStruct := &oc.Interface_Aggregation{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation", goStruct, true, true)
	if ok {
		return convertInterface_Aggregation_LagTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_LagTypePath) Get(t testing.TB) oc.E_IfAggregate_AggregationType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_LagTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfAggregate_AggregationType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfAggregate_AggregationType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Aggregation_LagTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type with a ONCE subscription.
func (n *Interface_Aggregation_LagTypePathAny) Get(t testing.TB) []oc.E_IfAggregate_AggregationType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfAggregate_AggregationType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type.
func (n *Interface_Aggregation_LagTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type in the given batch object.
func (n *Interface_Aggregation_LagTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type.
func (n *Interface_Aggregation_LagTypePath) Replace(t testing.TB, val oc.E_IfAggregate_AggregationType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type in the given batch object.
func (n *Interface_Aggregation_LagTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_IfAggregate_AggregationType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type.
func (n *Interface_Aggregation_LagTypePath) Update(t testing.TB, val oc.E_IfAggregate_AggregationType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/config/lag-type in the given batch object.
func (n *Interface_Aggregation_LagTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_IfAggregate_AggregationType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Aggregation_LagTypePath extracts the value of the leaf LagType from its parent oc.Interface_Aggregation
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfAggregate_AggregationType.
func convertInterface_Aggregation_LagTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Aggregation) *oc.QualifiedE_IfAggregate_AggregationType {
	t.Helper()
	qv := &oc.QualifiedE_IfAggregate_AggregationType{
		Metadata: md,
	}
	val := parent.LagType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_MinLinksPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Aggregation{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation", goStruct, true, true)
	if ok {
		return convertInterface_Aggregation_MinLinksPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_MinLinksPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_MinLinksPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Aggregation_MinLinksPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links with a ONCE subscription.
func (n *Interface_Aggregation_MinLinksPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links.
func (n *Interface_Aggregation_MinLinksPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links in the given batch object.
func (n *Interface_Aggregation_MinLinksPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links.
func (n *Interface_Aggregation_MinLinksPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links in the given batch object.
func (n *Interface_Aggregation_MinLinksPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links.
func (n *Interface_Aggregation_MinLinksPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/config/min-links in the given batch object.
func (n *Interface_Aggregation_MinLinksPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Aggregation_MinLinksPath extracts the value of the leaf MinLinks from its parent oc.Interface_Aggregation
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Aggregation_MinLinksPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Aggregation) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.MinLinks
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_SwitchedVlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_Aggregation_SwitchedVlan {
	t.Helper()
	goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation_SwitchedVlan", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Aggregation_SwitchedVlan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_SwitchedVlanPath) Get(t testing.TB) *oc.Interface_Aggregation_SwitchedVlan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_SwitchedVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Aggregation_SwitchedVlan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Aggregation_SwitchedVlan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation_SwitchedVlan", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Aggregation_SwitchedVlan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan with a ONCE subscription.
func (n *Interface_Aggregation_SwitchedVlanPathAny) Get(t testing.TB) []*oc.Interface_Aggregation_SwitchedVlan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Aggregation_SwitchedVlan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan.
func (n *Interface_Aggregation_SwitchedVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan.
func (n *Interface_Aggregation_SwitchedVlanPath) Replace(t testing.TB, val *oc.Interface_Aggregation_SwitchedVlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Aggregation_SwitchedVlan) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan.
func (n *Interface_Aggregation_SwitchedVlanPath) Update(t testing.TB, val *oc.Interface_Aggregation_SwitchedVlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Aggregation_SwitchedVlan) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Aggregation_SwitchedVlan_AccessVlanPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Aggregation_SwitchedVlan_AccessVlanPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan with a ONCE subscription.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/access-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_AccessVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Aggregation_SwitchedVlan_AccessVlanPath extracts the value of the leaf AccessVlan from its parent oc.Interface_Aggregation_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Aggregation_SwitchedVlan_AccessVlanPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Aggregation_SwitchedVlan) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AccessVlan
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Aggregation_SwitchedVlan_InterfaceModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) Get(t testing.TB) oc.E_VlanTypes_VlanModeType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_VlanModeType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Aggregation_SwitchedVlan_InterfaceModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode with a ONCE subscription.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePathAny) Get(t testing.TB) []oc.E_VlanTypes_VlanModeType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_VlanModeType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) Replace(t testing.TB, val oc.E_VlanTypes_VlanModeType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanModeType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) Update(t testing.TB, val oc.E_VlanTypes_VlanModeType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/interface-mode in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_InterfaceModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanModeType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Aggregation_SwitchedVlan_InterfaceModePath extracts the value of the leaf InterfaceMode from its parent oc.Interface_Aggregation_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_VlanModeType.
func convertInterface_Aggregation_SwitchedVlan_InterfaceModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Aggregation_SwitchedVlan) *oc.QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_VlanModeType{
		Metadata: md,
	}
	val := parent.InterfaceMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Aggregation_SwitchedVlan_NativeVlanPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Aggregation_SwitchedVlan_NativeVlanPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan with a ONCE subscription.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/native-vlan in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_NativeVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Aggregation_SwitchedVlan_NativeVlanPath extracts the value of the leaf NativeVlan from its parent oc.Interface_Aggregation_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Aggregation_SwitchedVlan_NativeVlanPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Aggregation_SwitchedVlan) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.NativeVlan
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) Lookup(t testing.TB) *oc.QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Aggregation_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Aggregation_SwitchedVlan_TrunkVlansPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) Get(t testing.TB) []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Aggregation_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Aggregation_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Aggregation_SwitchedVlan_TrunkVlansPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans with a ONCE subscription.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPathAny) Get(t testing.TB) [][]oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) Replace(t testing.TB, val []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) Update(t testing.TB, val []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/aggregation/switched-vlan/config/trunk-vlans in the given batch object.
func (n *Interface_Aggregation_SwitchedVlan_TrunkVlansPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Aggregation_SwitchedVlan_TrunkVlansPath extracts the value of the leaf TrunkVlans from its parent oc.Interface_Aggregation_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice.
func convertInterface_Aggregation_SwitchedVlan_TrunkVlansPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Aggregation_SwitchedVlan) *oc.QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice{
		Metadata: md,
	}
	val := parent.TrunkVlans
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/description with a ONCE subscription.
func (n *Interface_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/description.
func (n *Interface_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/description in the given batch object.
func (n *Interface_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/description.
func (n *Interface_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/description in the given batch object.
func (n *Interface_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/description.
func (n *Interface_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/description in the given batch object.
func (n *Interface_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_DescriptionPath extracts the value of the leaf Description from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/enabled with a ONCE subscription.
func (n *Interface_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/enabled.
func (n *Interface_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/enabled in the given batch object.
func (n *Interface_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/enabled.
func (n *Interface_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/enabled in the given batch object.
func (n *Interface_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/enabled.
func (n *Interface_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/enabled in the given batch object.
func (n *Interface_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_EthernetPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ethernet {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Ethernet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_EthernetPath) Get(t testing.TB) *oc.Interface_Ethernet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_EthernetPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ethernet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ethernet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Ethernet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet with a ONCE subscription.
func (n *Interface_EthernetPathAny) Get(t testing.TB) []*oc.Interface_Ethernet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Ethernet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet.
func (n *Interface_EthernetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet in the given batch object.
func (n *Interface_EthernetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet.
func (n *Interface_EthernetPath) Replace(t testing.TB, val *oc.Interface_Ethernet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet in the given batch object.
func (n *Interface_EthernetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Ethernet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet.
func (n *Interface_EthernetPath) Update(t testing.TB, val *oc.Interface_Ethernet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet in the given batch object.
func (n *Interface_EthernetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Ethernet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_AggregateIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_AggregateIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_AggregateIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_AggregateIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_AggregateIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id with a ONCE subscription.
func (n *Interface_Ethernet_AggregateIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id.
func (n *Interface_Ethernet_AggregateIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id in the given batch object.
func (n *Interface_Ethernet_AggregateIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id.
func (n *Interface_Ethernet_AggregateIdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id in the given batch object.
func (n *Interface_Ethernet_AggregateIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id.
func (n *Interface_Ethernet_AggregateIdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/aggregate-id in the given batch object.
func (n *Interface_Ethernet_AggregateIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_AggregateIdPath extracts the value of the leaf AggregateId from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ethernet_AggregateIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AggregateId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_AutoNegotiatePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_AutoNegotiatePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAutoNegotiate())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_AutoNegotiatePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_AutoNegotiatePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_AutoNegotiatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate with a ONCE subscription.
func (n *Interface_Ethernet_AutoNegotiatePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate.
func (n *Interface_Ethernet_AutoNegotiatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate in the given batch object.
func (n *Interface_Ethernet_AutoNegotiatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate.
func (n *Interface_Ethernet_AutoNegotiatePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate in the given batch object.
func (n *Interface_Ethernet_AutoNegotiatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate.
func (n *Interface_Ethernet_AutoNegotiatePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/auto-negotiate in the given batch object.
func (n *Interface_Ethernet_AutoNegotiatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_AutoNegotiatePath extracts the value of the leaf AutoNegotiate from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Ethernet_AutoNegotiatePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AutoNegotiate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_DuplexModePath) Lookup(t testing.TB) *oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_DuplexModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_DuplexModePath) Get(t testing.TB) oc.E_Ethernet_DuplexMode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_DuplexModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ethernet_DuplexMode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_DuplexModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode with a ONCE subscription.
func (n *Interface_Ethernet_DuplexModePathAny) Get(t testing.TB) []oc.E_Ethernet_DuplexMode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ethernet_DuplexMode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode.
func (n *Interface_Ethernet_DuplexModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode in the given batch object.
func (n *Interface_Ethernet_DuplexModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode.
func (n *Interface_Ethernet_DuplexModePath) Replace(t testing.TB, val oc.E_Ethernet_DuplexMode) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode in the given batch object.
func (n *Interface_Ethernet_DuplexModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Ethernet_DuplexMode) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode.
func (n *Interface_Ethernet_DuplexModePath) Update(t testing.TB, val oc.E_Ethernet_DuplexMode) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/duplex-mode in the given batch object.
func (n *Interface_Ethernet_DuplexModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Ethernet_DuplexMode) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Ethernet_DuplexModePath extracts the value of the leaf DuplexMode from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ethernet_DuplexMode.
func convertInterface_Ethernet_DuplexModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	qv := &oc.QualifiedE_Ethernet_DuplexMode{
		Metadata: md,
	}
	val := parent.DuplexMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_EnableFlowControlPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_EnableFlowControlPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableFlowControl())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_EnableFlowControlPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_EnableFlowControlPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control with a ONCE subscription.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control.
func (n *Interface_Ethernet_EnableFlowControlPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control in the given batch object.
func (n *Interface_Ethernet_EnableFlowControlPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control.
func (n *Interface_Ethernet_EnableFlowControlPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control in the given batch object.
func (n *Interface_Ethernet_EnableFlowControlPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control.
func (n *Interface_Ethernet_EnableFlowControlPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/enable-flow-control in the given batch object.
func (n *Interface_Ethernet_EnableFlowControlPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_EnableFlowControlPath extracts the value of the leaf EnableFlowControl from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Ethernet_EnableFlowControlPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableFlowControl
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_FecModePath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_FecModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_FecModePath) Get(t testing.TB) oc.E_IfEthernet_INTERFACE_FEC {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_FecModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_INTERFACE_FEC
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_FecModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode with a ONCE subscription.
func (n *Interface_Ethernet_FecModePathAny) Get(t testing.TB) []oc.E_IfEthernet_INTERFACE_FEC {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_INTERFACE_FEC
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode.
func (n *Interface_Ethernet_FecModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode in the given batch object.
func (n *Interface_Ethernet_FecModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode.
func (n *Interface_Ethernet_FecModePath) Replace(t testing.TB, val oc.E_IfEthernet_INTERFACE_FEC) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode in the given batch object.
func (n *Interface_Ethernet_FecModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_INTERFACE_FEC) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode.
func (n *Interface_Ethernet_FecModePath) Update(t testing.TB, val oc.E_IfEthernet_INTERFACE_FEC) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/fec-mode in the given batch object.
func (n *Interface_Ethernet_FecModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_INTERFACE_FEC) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Ethernet_FecModePath extracts the value of the leaf FecMode from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_INTERFACE_FEC.
func convertInterface_Ethernet_FecModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_INTERFACE_FEC{
		Metadata: md,
	}
	val := parent.FecMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_MacAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_MacAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_MacAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_MacAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_MacAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address with a ONCE subscription.
func (n *Interface_Ethernet_MacAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address.
func (n *Interface_Ethernet_MacAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address in the given batch object.
func (n *Interface_Ethernet_MacAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address.
func (n *Interface_Ethernet_MacAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address in the given batch object.
func (n *Interface_Ethernet_MacAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address.
func (n *Interface_Ethernet_MacAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/mac-address in the given batch object.
func (n *Interface_Ethernet_MacAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_MacAddressPath extracts the value of the leaf MacAddress from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ethernet_MacAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MacAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_PortSpeedPath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_PortSpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_PortSpeedPath) Get(t testing.TB) oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_PortSpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_PortSpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed with a ONCE subscription.
func (n *Interface_Ethernet_PortSpeedPathAny) Get(t testing.TB) []oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_ETHERNET_SPEED
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed.
func (n *Interface_Ethernet_PortSpeedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed in the given batch object.
func (n *Interface_Ethernet_PortSpeedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed.
func (n *Interface_Ethernet_PortSpeedPath) Replace(t testing.TB, val oc.E_IfEthernet_ETHERNET_SPEED) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed in the given batch object.
func (n *Interface_Ethernet_PortSpeedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_ETHERNET_SPEED) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed.
func (n *Interface_Ethernet_PortSpeedPath) Update(t testing.TB, val oc.E_IfEthernet_ETHERNET_SPEED) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/port-speed in the given batch object.
func (n *Interface_Ethernet_PortSpeedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_ETHERNET_SPEED) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Ethernet_PortSpeedPath extracts the value of the leaf PortSpeed from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_ETHERNET_SPEED.
func convertInterface_Ethernet_PortSpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_ETHERNET_SPEED{
		Metadata: md,
	}
	val := parent.PortSpeed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_StandaloneLinkTrainingPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetStandaloneLinkTraining())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_StandaloneLinkTrainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training with a ONCE subscription.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training in the given batch object.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training in the given batch object.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/config/standalone-link-training in the given batch object.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_StandaloneLinkTrainingPath extracts the value of the leaf StandaloneLinkTraining from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Ethernet_StandaloneLinkTrainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.StandaloneLinkTraining
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_SwitchedVlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ethernet_SwitchedVlan {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_SwitchedVlan", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Ethernet_SwitchedVlan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_SwitchedVlanPath) Get(t testing.TB) *oc.Interface_Ethernet_SwitchedVlan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_SwitchedVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ethernet_SwitchedVlan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ethernet_SwitchedVlan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_SwitchedVlan", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Ethernet_SwitchedVlan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan with a ONCE subscription.
func (n *Interface_Ethernet_SwitchedVlanPathAny) Get(t testing.TB) []*oc.Interface_Ethernet_SwitchedVlan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Ethernet_SwitchedVlan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan.
func (n *Interface_Ethernet_SwitchedVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan.
func (n *Interface_Ethernet_SwitchedVlanPath) Replace(t testing.TB, val *oc.Interface_Ethernet_SwitchedVlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Ethernet_SwitchedVlan) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan.
func (n *Interface_Ethernet_SwitchedVlanPath) Update(t testing.TB, val *oc.Interface_Ethernet_SwitchedVlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Ethernet_SwitchedVlan) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_SwitchedVlan_AccessVlanPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_SwitchedVlan_AccessVlanPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan with a ONCE subscription.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/access-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_AccessVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_SwitchedVlan_AccessVlanPath extracts the value of the leaf AccessVlan from its parent oc.Interface_Ethernet_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Ethernet_SwitchedVlan_AccessVlanPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_SwitchedVlan) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AccessVlan
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_SwitchedVlan_InterfaceModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) Get(t testing.TB) oc.E_VlanTypes_VlanModeType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_VlanModeType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_SwitchedVlan_InterfaceModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode with a ONCE subscription.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePathAny) Get(t testing.TB) []oc.E_VlanTypes_VlanModeType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_VlanModeType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) Replace(t testing.TB, val oc.E_VlanTypes_VlanModeType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanModeType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) Update(t testing.TB, val oc.E_VlanTypes_VlanModeType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/interface-mode in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_InterfaceModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_VlanModeType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Ethernet_SwitchedVlan_InterfaceModePath extracts the value of the leaf InterfaceMode from its parent oc.Interface_Ethernet_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_VlanModeType.
func convertInterface_Ethernet_SwitchedVlan_InterfaceModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_SwitchedVlan) *oc.QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_VlanModeType{
		Metadata: md,
	}
	val := parent.InterfaceMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_SwitchedVlan_NativeVlanPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_SwitchedVlan_NativeVlanPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan with a ONCE subscription.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/native-vlan in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_NativeVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Ethernet_SwitchedVlan_NativeVlanPath extracts the value of the leaf NativeVlan from its parent oc.Interface_Ethernet_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Ethernet_SwitchedVlan_NativeVlanPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_SwitchedVlan) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.NativeVlan
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_SwitchedVlan", goStruct, true, true)
	if ok {
		return convertInterface_Ethernet_SwitchedVlan_TrunkVlansPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) Get(t testing.TB) []oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_SwitchedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_SwitchedVlan", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_SwitchedVlan_TrunkVlansPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans with a ONCE subscription.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPathAny) Get(t testing.TB) [][]oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) Replace(t testing.TB, val []oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) Update(t testing.TB, val []oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/ethernet/switched-vlan/config/trunk-vlans in the given batch object.
func (n *Interface_Ethernet_SwitchedVlan_TrunkVlansPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.Interface_Ethernet_SwitchedVlan_TrunkVlans_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Ethernet_SwitchedVlan_TrunkVlansPath extracts the value of the leaf TrunkVlans from its parent oc.Interface_Ethernet_SwitchedVlan
// and combines the update with an existing Metadata to return a *oc.QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice.
func convertInterface_Ethernet_SwitchedVlan_TrunkVlansPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_SwitchedVlan) *oc.QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice{
		Metadata: md,
	}
	val := parent.TrunkVlans
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HoldTimePath) Lookup(t testing.TB) *oc.QualifiedInterface_HoldTime {
	t.Helper()
	goStruct := &oc.Interface_HoldTime{}
	md, ok := oc.Lookup(t, n, "Interface_HoldTime", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_HoldTime{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HoldTimePath) Get(t testing.TB) *oc.Interface_HoldTime {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HoldTimePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_HoldTime {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_HoldTime
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_HoldTime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_HoldTime", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_HoldTime{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription.
func (n *Interface_HoldTimePathAny) Get(t testing.TB) []*oc.Interface_HoldTime {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_HoldTime
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/hold-time.
func (n *Interface_HoldTimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/hold-time in the given batch object.
func (n *Interface_HoldTimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/hold-time.
func (n *Interface_HoldTimePath) Replace(t testing.TB, val *oc.Interface_HoldTime) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/hold-time in the given batch object.
func (n *Interface_HoldTimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_HoldTime) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/hold-time.
func (n *Interface_HoldTimePath) Update(t testing.TB, val *oc.Interface_HoldTime) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/hold-time in the given batch object.
func (n *Interface_HoldTimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_HoldTime) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/config/down with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HoldTime_DownPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_HoldTime{}
	md, ok := oc.Lookup(t, n, "Interface_HoldTime", goStruct, true, true)
	if ok {
		return convertInterface_HoldTime_DownPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetDown())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/config/down with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HoldTime_DownPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/config/down with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HoldTime_DownPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_HoldTime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_HoldTime", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_HoldTime_DownPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/config/down with a ONCE subscription.
func (n *Interface_HoldTime_DownPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/hold-time/config/down.
func (n *Interface_HoldTime_DownPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/hold-time/config/down in the given batch object.
func (n *Interface_HoldTime_DownPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/hold-time/config/down.
func (n *Interface_HoldTime_DownPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/hold-time/config/down in the given batch object.
func (n *Interface_HoldTime_DownPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/hold-time/config/down.
func (n *Interface_HoldTime_DownPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/hold-time/config/down in the given batch object.
func (n *Interface_HoldTime_DownPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_HoldTime_DownPath extracts the value of the leaf Down from its parent oc.Interface_HoldTime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_HoldTime_DownPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_HoldTime) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Down
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/config/up with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HoldTime_UpPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_HoldTime{}
	md, ok := oc.Lookup(t, n, "Interface_HoldTime", goStruct, true, true)
	if ok {
		return convertInterface_HoldTime_UpPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetUp())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/config/up with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HoldTime_UpPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/config/up with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HoldTime_UpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_HoldTime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_HoldTime", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_HoldTime_UpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/config/up with a ONCE subscription.
func (n *Interface_HoldTime_UpPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/hold-time/config/up.
func (n *Interface_HoldTime_UpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/hold-time/config/up in the given batch object.
func (n *Interface_HoldTime_UpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/hold-time/config/up.
func (n *Interface_HoldTime_UpPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/hold-time/config/up in the given batch object.
func (n *Interface_HoldTime_UpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/hold-time/config/up.
func (n *Interface_HoldTime_UpPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/hold-time/config/up in the given batch object.
func (n *Interface_HoldTime_UpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_HoldTime_UpPath extracts the value of the leaf Up from its parent oc.Interface_HoldTime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_HoldTime_UpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_HoldTime) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Up
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/loopback-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_LoopbackModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_LoopbackModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetLoopbackMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/loopback-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_LoopbackModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/loopback-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_LoopbackModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_LoopbackModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/loopback-mode with a ONCE subscription.
func (n *Interface_LoopbackModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/loopback-mode.
func (n *Interface_LoopbackModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/loopback-mode in the given batch object.
func (n *Interface_LoopbackModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/loopback-mode.
func (n *Interface_LoopbackModePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/loopback-mode in the given batch object.
func (n *Interface_LoopbackModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/loopback-mode.
func (n *Interface_LoopbackModePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/loopback-mode in the given batch object.
func (n *Interface_LoopbackModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_LoopbackModePath extracts the value of the leaf LoopbackMode from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_LoopbackModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.LoopbackMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/mtu with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_MtuPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/mtu with a ONCE subscription.
func (n *Interface_MtuPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/mtu.
func (n *Interface_MtuPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/mtu in the given batch object.
func (n *Interface_MtuPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/mtu.
func (n *Interface_MtuPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/mtu in the given batch object.
func (n *Interface_MtuPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/mtu.
func (n *Interface_MtuPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/mtu in the given batch object.
func (n *Interface_MtuPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, true)
	if ok {
		return convertInterface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/config/name with a ONCE subscription.
func (n *Interface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/config/name.
func (n *Interface_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/config/name in the given batch object.
func (n *Interface_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/config/name.
func (n *Interface_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/config/name in the given batch object.
func (n *Interface_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/config/name.
func (n *Interface_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/config/name in the given batch object.
func (n *Interface_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_NamePath extracts the value of the leaf Name from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlanPath) Get(t testing.TB) *oc.Interface_RoutedVlan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription.
func (n *Interface_RoutedVlanPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan.
func (n *Interface_RoutedVlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan in the given batch object.
func (n *Interface_RoutedVlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan.
func (n *Interface_RoutedVlanPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan in the given batch object.
func (n *Interface_RoutedVlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan.
func (n *Interface_RoutedVlanPath) Update(t testing.TB, val *oc.Interface_RoutedVlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan in the given batch object.
func (n *Interface_RoutedVlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4Path) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4PathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4.
func (n *Interface_RoutedVlan_Ipv4Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 in the given batch object.
func (n *Interface_RoutedVlan_Ipv4Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4.
func (n *Interface_RoutedVlan_Ipv4Path) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 in the given batch object.
func (n *Interface_RoutedVlan_Ipv4Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4.
func (n *Interface_RoutedVlan_Ipv4Path) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 in the given batch object.
func (n *Interface_RoutedVlan_Ipv4Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Address) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Address) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Address) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Address) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/ip in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_RoutedVlan_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAcceptMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
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
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath extracts the value of the leaf AcceptMode from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAdvertisementInterval())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
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
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath extracts the value of the leaf AdvertisementInterval from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Replace(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Update(t testing.TB, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

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
// failing the test fatally is no value is present at the path.
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
