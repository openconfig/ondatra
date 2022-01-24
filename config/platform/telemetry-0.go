package platform

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
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

// Lookup fetches the value at /openconfig-platform/components/component/fabric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_FabricPath) Lookup(t testing.TB) *oc.QualifiedComponent_Fabric {
	t.Helper()
	goStruct := &oc.Component_Fabric{}
	md, ok := oc.Lookup(t, n, "Component_Fabric", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Fabric{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/fabric with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_FabricPath) Get(t testing.TB) *oc.Component_Fabric {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/fabric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_FabricPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Fabric {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Fabric
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Fabric{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Fabric", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Fabric{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/fabric with a ONCE subscription.
func (n *Component_FabricPathAny) Get(t testing.TB) []*oc.Component_Fabric {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Fabric
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/fabric.
func (n *Component_FabricPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/fabric in the given batch object.
func (n *Component_FabricPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/fabric.
func (n *Component_FabricPath) Replace(t testing.TB, val *oc.Component_Fabric) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/fabric in the given batch object.
func (n *Component_FabricPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Fabric) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/fabric.
func (n *Component_FabricPath) Update(t testing.TB, val *oc.Component_Fabric) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/fabric in the given batch object.
func (n *Component_FabricPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Fabric) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/fan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_FanPath) Lookup(t testing.TB) *oc.QualifiedComponent_Fan {
	t.Helper()
	goStruct := &oc.Component_Fan{}
	md, ok := oc.Lookup(t, n, "Component_Fan", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Fan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/fan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_FanPath) Get(t testing.TB) *oc.Component_Fan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/fan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_FanPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Fan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Fan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Fan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Fan", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Fan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/fan with a ONCE subscription.
func (n *Component_FanPathAny) Get(t testing.TB) []*oc.Component_Fan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Fan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/fan.
func (n *Component_FanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/fan in the given batch object.
func (n *Component_FanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/fan.
func (n *Component_FanPath) Replace(t testing.TB, val *oc.Component_Fan) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/fan in the given batch object.
func (n *Component_FanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Fan) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/fan.
func (n *Component_FanPath) Update(t testing.TB, val *oc.Component_Fan) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/fan in the given batch object.
func (n *Component_FanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Fan) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuitPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuitPath) Get(t testing.TB) *oc.Component_IntegratedCircuit {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuitPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription.
func (n *Component_IntegratedCircuitPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit.
func (n *Component_IntegratedCircuitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit in the given batch object.
func (n *Component_IntegratedCircuitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit.
func (n *Component_IntegratedCircuitPath) Replace(t testing.TB, val *oc.Component_IntegratedCircuit) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit in the given batch object.
func (n *Component_IntegratedCircuitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit.
func (n *Component_IntegratedCircuitPath) Update(t testing.TB, val *oc.Component_IntegratedCircuit) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit in the given batch object.
func (n *Component_IntegratedCircuitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
