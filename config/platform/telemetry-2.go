package platform

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

// Lookup fetches the value at /openconfig-platform/components/component/power-supply with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PowerSupplyPath) Lookup(t testing.TB) *oc.QualifiedComponent_PowerSupply {
	t.Helper()
	goStruct := &oc.Component_PowerSupply{}
	md, ok := oc.Lookup(t, n, "Component_PowerSupply", goStruct, false, true)
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_PowerSupply", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/power-supply.
func (n *Component_PowerSupplyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/power-supply in the given batch object.
func (n *Component_PowerSupplyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/power-supply.
func (n *Component_PowerSupplyPath) Replace(t testing.TB, val *oc.Component_PowerSupply) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/power-supply in the given batch object.
func (n *Component_PowerSupplyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_PowerSupply) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/power-supply.
func (n *Component_PowerSupplyPath) Update(t testing.TB, val *oc.Component_PowerSupply) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/power-supply in the given batch object.
func (n *Component_PowerSupplyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_PowerSupply) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PropertyPath) Lookup(t testing.TB) *oc.QualifiedComponent_Property {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, false, true)
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/properties/property.
func (n *Component_PropertyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/properties/property in the given batch object.
func (n *Component_PropertyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/properties/property.
func (n *Component_PropertyPath) Replace(t testing.TB, val *oc.Component_Property) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/properties/property in the given batch object.
func (n *Component_PropertyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Property) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/properties/property.
func (n *Component_PropertyPath) Update(t testing.TB, val *oc.Component_Property) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/properties/property in the given batch object.
func (n *Component_PropertyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Property) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, true)
	if ok {
		return convertComponent_Property_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Property_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription.
func (n *Component_Property_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/properties/property/config/name.
func (n *Component_Property_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/properties/property/config/name in the given batch object.
func (n *Component_Property_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/properties/property/config/name.
func (n *Component_Property_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/properties/property/config/name in the given batch object.
func (n *Component_Property_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/properties/property/config/name.
func (n *Component_Property_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/properties/property/config/name in the given batch object.
func (n *Component_Property_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_ValuePath) Lookup(t testing.TB) *oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, true)
	if ok {
		return convertComponent_Property_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_ValuePath) Get(t testing.TB) oc.Component_Property_Value_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Property_Value_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Property_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription.
func (n *Component_Property_ValuePathAny) Get(t testing.TB) []oc.Component_Property_Value_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Component_Property_Value_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/properties/property/config/value.
func (n *Component_Property_ValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/properties/property/config/value in the given batch object.
func (n *Component_Property_ValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/properties/property/config/value.
func (n *Component_Property_ValuePath) Replace(t testing.TB, val oc.Component_Property_Value_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/properties/property/config/value in the given batch object.
func (n *Component_Property_ValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Component_Property_Value_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/properties/property/config/value.
func (n *Component_Property_ValuePath) Update(t testing.TB, val oc.Component_Property_Value_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/properties/property/config/value in the given batch object.
func (n *Component_Property_ValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Component_Property_Value_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
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

// Lookup fetches the value at /openconfig-platform/components/component/software-module with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SoftwareModulePath) Lookup(t testing.TB) *oc.QualifiedComponent_SoftwareModule {
	t.Helper()
	goStruct := &oc.Component_SoftwareModule{}
	md, ok := oc.Lookup(t, n, "Component_SoftwareModule", goStruct, false, true)
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_SoftwareModule", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/software-module.
func (n *Component_SoftwareModulePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/software-module in the given batch object.
func (n *Component_SoftwareModulePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/software-module.
func (n *Component_SoftwareModulePath) Replace(t testing.TB, val *oc.Component_SoftwareModule) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/software-module in the given batch object.
func (n *Component_SoftwareModulePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_SoftwareModule) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/software-module.
func (n *Component_SoftwareModulePath) Update(t testing.TB, val *oc.Component_SoftwareModule) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/software-module in the given batch object.
func (n *Component_SoftwareModulePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_SoftwareModule) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/storage with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_StoragePath) Lookup(t testing.TB) *oc.QualifiedComponent_Storage {
	t.Helper()
	goStruct := &oc.Component_Storage{}
	md, ok := oc.Lookup(t, n, "Component_Storage", goStruct, false, true)
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Storage", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/storage.
func (n *Component_StoragePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/storage in the given batch object.
func (n *Component_StoragePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/storage.
func (n *Component_StoragePath) Replace(t testing.TB, val *oc.Component_Storage) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/storage in the given batch object.
func (n *Component_StoragePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Storage) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/storage.
func (n *Component_StoragePath) Update(t testing.TB, val *oc.Component_Storage) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/storage in the given batch object.
func (n *Component_StoragePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Storage) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SubcomponentPath) Lookup(t testing.TB) *oc.QualifiedComponent_Subcomponent {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, false, true)
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/subcomponents/subcomponent.
func (n *Component_SubcomponentPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/subcomponents/subcomponent in the given batch object.
func (n *Component_SubcomponentPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/subcomponents/subcomponent.
func (n *Component_SubcomponentPath) Replace(t testing.TB, val *oc.Component_Subcomponent) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/subcomponents/subcomponent in the given batch object.
func (n *Component_SubcomponentPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Subcomponent) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/subcomponents/subcomponent.
func (n *Component_SubcomponentPath) Update(t testing.TB, val *oc.Component_Subcomponent) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/subcomponents/subcomponent in the given batch object.
func (n *Component_SubcomponentPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Subcomponent) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Subcomponent_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, true, true)
	if ok {
		return convertComponent_Subcomponent_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Subcomponent_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Subcomponent_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Subcomponent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Subcomponent_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription.
func (n *Component_Subcomponent_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/subcomponents/subcomponent/config/name.
func (n *Component_Subcomponent_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/subcomponents/subcomponent/config/name in the given batch object.
func (n *Component_Subcomponent_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/subcomponents/subcomponent/config/name.
func (n *Component_Subcomponent_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/subcomponents/subcomponent/config/name in the given batch object.
func (n *Component_Subcomponent_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/subcomponents/subcomponent/config/name.
func (n *Component_Subcomponent_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/subcomponents/subcomponent/config/name in the given batch object.
func (n *Component_Subcomponent_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
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
