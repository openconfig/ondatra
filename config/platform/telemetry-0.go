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

// Lookup fetches the value at /openconfig-platform/components/component with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *ComponentPath) Lookup(t testing.TB) *oc.QualifiedComponent {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component.
func (n *ComponentPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component in the given batch object.
func (n *ComponentPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component.
func (n *ComponentPath) Replace(t testing.TB, val *oc.Component) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component in the given batch object.
func (n *ComponentPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component.
func (n *ComponentPath) Update(t testing.TB, val *oc.Component) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component in the given batch object.
func (n *ComponentPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/backplane with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_BackplanePath) Lookup(t testing.TB) *oc.QualifiedComponent_Backplane {
	t.Helper()
	goStruct := &oc.Component_Backplane{}
	md, ok := oc.Lookup(t, n, "Component_Backplane", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Backplane{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/backplane with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Backplane", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/backplane.
func (n *Component_BackplanePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/backplane in the given batch object.
func (n *Component_BackplanePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/backplane.
func (n *Component_BackplanePath) Replace(t testing.TB, val *oc.Component_Backplane) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/backplane in the given batch object.
func (n *Component_BackplanePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Backplane) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/backplane.
func (n *Component_BackplanePath) Update(t testing.TB, val *oc.Component_Backplane) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/backplane in the given batch object.
func (n *Component_BackplanePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Backplane) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/chassis with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_ChassisPath) Lookup(t testing.TB) *oc.QualifiedComponent_Chassis {
	t.Helper()
	goStruct := &oc.Component_Chassis{}
	md, ok := oc.Lookup(t, n, "Component_Chassis", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Chassis{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/chassis with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Chassis", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/chassis.
func (n *Component_ChassisPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/chassis in the given batch object.
func (n *Component_ChassisPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/chassis.
func (n *Component_ChassisPath) Replace(t testing.TB, val *oc.Component_Chassis) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/chassis in the given batch object.
func (n *Component_ChassisPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Chassis) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/chassis.
func (n *Component_ChassisPath) Update(t testing.TB, val *oc.Component_Chassis) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/chassis in the given batch object.
func (n *Component_ChassisPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Chassis) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/chassis/config/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Chassis_IdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Component_Chassis{}
	md, ok := oc.Lookup(t, n, "Component_Chassis", goStruct, true, true)
	if ok {
		return convertComponent_Chassis_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/chassis/config/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Chassis_IdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/chassis/config/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Chassis_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Chassis{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Chassis", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Chassis_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/chassis/config/id with a ONCE subscription.
func (n *Component_Chassis_IdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/chassis/config/id.
func (n *Component_Chassis_IdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/chassis/config/id in the given batch object.
func (n *Component_Chassis_IdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/chassis/config/id.
func (n *Component_Chassis_IdPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/chassis/config/id in the given batch object.
func (n *Component_Chassis_IdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/chassis/config/id.
func (n *Component_Chassis_IdPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/chassis/config/id in the given batch object.
func (n *Component_Chassis_IdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Chassis_IdPath extracts the value of the leaf Id from its parent oc.Component_Chassis
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertComponent_Chassis_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Chassis) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/chassis/utilization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Chassis_UtilizationPath) Lookup(t testing.TB) *oc.QualifiedComponent_Chassis_Utilization {
	t.Helper()
	goStruct := &oc.Component_Chassis_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Chassis_Utilization", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Chassis_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/chassis/utilization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Chassis_UtilizationPath) Get(t testing.TB) *oc.Component_Chassis_Utilization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/chassis/utilization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Chassis_UtilizationPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Chassis_Utilization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Chassis_Utilization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Chassis_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Chassis_Utilization", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Chassis_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/chassis/utilization with a ONCE subscription.
func (n *Component_Chassis_UtilizationPathAny) Get(t testing.TB) []*oc.Component_Chassis_Utilization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Chassis_Utilization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/chassis/utilization.
func (n *Component_Chassis_UtilizationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/chassis/utilization in the given batch object.
func (n *Component_Chassis_UtilizationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/chassis/utilization.
func (n *Component_Chassis_UtilizationPath) Replace(t testing.TB, val *oc.Component_Chassis_Utilization) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/chassis/utilization in the given batch object.
func (n *Component_Chassis_UtilizationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Chassis_Utilization) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/chassis/utilization.
func (n *Component_Chassis_UtilizationPath) Update(t testing.TB, val *oc.Component_Chassis_Utilization) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/chassis/utilization in the given batch object.
func (n *Component_Chassis_UtilizationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Chassis_Utilization) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/chassis/utilization/resources/resource with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Chassis_Utilization_ResourcePath) Lookup(t testing.TB) *oc.QualifiedComponent_Chassis_Utilization_Resource {
	t.Helper()
	goStruct := &oc.Component_Chassis_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_Chassis_Utilization_Resource", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Chassis_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/chassis/utilization/resources/resource with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Chassis_Utilization_ResourcePath) Get(t testing.TB) *oc.Component_Chassis_Utilization_Resource {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/chassis/utilization/resources/resource with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Chassis_Utilization_ResourcePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Chassis_Utilization_Resource {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Chassis_Utilization_Resource
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Chassis_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Chassis_Utilization_Resource", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Chassis_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/chassis/utilization/resources/resource with a ONCE subscription.
func (n *Component_Chassis_Utilization_ResourcePathAny) Get(t testing.TB) []*oc.Component_Chassis_Utilization_Resource {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Chassis_Utilization_Resource
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/chassis/utilization/resources/resource.
func (n *Component_Chassis_Utilization_ResourcePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/chassis/utilization/resources/resource in the given batch object.
func (n *Component_Chassis_Utilization_ResourcePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/chassis/utilization/resources/resource.
func (n *Component_Chassis_Utilization_ResourcePath) Replace(t testing.TB, val *oc.Component_Chassis_Utilization_Resource) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/chassis/utilization/resources/resource in the given batch object.
func (n *Component_Chassis_Utilization_ResourcePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Chassis_Utilization_Resource) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/chassis/utilization/resources/resource.
func (n *Component_Chassis_Utilization_ResourcePath) Update(t testing.TB, val *oc.Component_Chassis_Utilization_Resource) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/chassis/utilization/resources/resource in the given batch object.
func (n *Component_Chassis_Utilization_ResourcePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Chassis_Utilization_Resource) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Chassis_Utilization_Resource_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Chassis_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_Chassis_Utilization_Resource", goStruct, true, true)
	if ok {
		return convertComponent_Chassis_Utilization_Resource_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Chassis_Utilization_Resource_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Chassis_Utilization_Resource_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Chassis_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Chassis_Utilization_Resource", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Chassis_Utilization_Resource_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name with a ONCE subscription.
func (n *Component_Chassis_Utilization_Resource_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name.
func (n *Component_Chassis_Utilization_Resource_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name in the given batch object.
func (n *Component_Chassis_Utilization_Resource_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name.
func (n *Component_Chassis_Utilization_Resource_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name in the given batch object.
func (n *Component_Chassis_Utilization_Resource_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name.
func (n *Component_Chassis_Utilization_Resource_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/chassis/utilization/resources/resource/config/name in the given batch object.
func (n *Component_Chassis_Utilization_Resource_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Chassis_Utilization_Resource_NamePath extracts the value of the leaf Name from its parent oc.Component_Chassis_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Chassis_Utilization_Resource_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Chassis_Utilization_Resource) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/controller-card with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_ControllerCardPath) Lookup(t testing.TB) *oc.QualifiedComponent_ControllerCard {
	t.Helper()
	goStruct := &oc.Component_ControllerCard{}
	md, ok := oc.Lookup(t, n, "Component_ControllerCard", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_ControllerCard{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/controller-card with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_ControllerCardPath) Get(t testing.TB) *oc.Component_ControllerCard {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/controller-card with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_ControllerCardPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_ControllerCard {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_ControllerCard
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_ControllerCard{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_ControllerCard", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_ControllerCard{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/controller-card with a ONCE subscription.
func (n *Component_ControllerCardPathAny) Get(t testing.TB) []*oc.Component_ControllerCard {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_ControllerCard
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/controller-card.
func (n *Component_ControllerCardPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/controller-card in the given batch object.
func (n *Component_ControllerCardPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/controller-card.
func (n *Component_ControllerCardPath) Replace(t testing.TB, val *oc.Component_ControllerCard) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/controller-card in the given batch object.
func (n *Component_ControllerCardPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_ControllerCard) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/controller-card.
func (n *Component_ControllerCardPath) Update(t testing.TB, val *oc.Component_ControllerCard) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/controller-card in the given batch object.
func (n *Component_ControllerCardPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_ControllerCard) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_CpuPath) Lookup(t testing.TB) *oc.QualifiedComponent_Cpu {
	t.Helper()
	goStruct := &oc.Component_Cpu{}
	md, ok := oc.Lookup(t, n, "Component_Cpu", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Cpu{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/cpu.
func (n *Component_CpuPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/cpu in the given batch object.
func (n *Component_CpuPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/cpu.
func (n *Component_CpuPath) Replace(t testing.TB, val *oc.Component_Cpu) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/cpu in the given batch object.
func (n *Component_CpuPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Cpu) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/cpu.
func (n *Component_CpuPath) Update(t testing.TB, val *oc.Component_Cpu) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/cpu in the given batch object.
func (n *Component_CpuPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Cpu) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_UtilizationPath) Lookup(t testing.TB) *oc.QualifiedComponent_Cpu_Utilization {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Cpu_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, false, true)
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

// Delete deletes the configuration at /openconfig-platform/components/component/cpu/utilization.
func (n *Component_Cpu_UtilizationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/cpu/utilization in the given batch object.
func (n *Component_Cpu_UtilizationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/cpu/utilization.
func (n *Component_Cpu_UtilizationPath) Replace(t testing.TB, val *oc.Component_Cpu_Utilization) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/cpu/utilization in the given batch object.
func (n *Component_Cpu_UtilizationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Cpu_Utilization) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/cpu/utilization.
func (n *Component_Cpu_UtilizationPath) Update(t testing.TB, val *oc.Component_Cpu_Utilization) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/cpu/utilization in the given batch object.
func (n *Component_Cpu_UtilizationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Cpu_Utilization) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
