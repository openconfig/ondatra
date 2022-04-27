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

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_BackplaneFacingCapacity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity in the given batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Replace(t testing.TB, val *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity in the given batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Update(t testing.TB, val *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity in the given batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_MemoryPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Memory {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Memory{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Memory", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Memory{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_MemoryPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Memory {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_MemoryPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Memory {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Memory
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Memory", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Memory{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription.
func (n *Component_IntegratedCircuit_MemoryPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Memory {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Memory
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit/memory.
func (n *Component_IntegratedCircuit_MemoryPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit/memory in the given batch object.
func (n *Component_IntegratedCircuit_MemoryPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit/memory.
func (n *Component_IntegratedCircuit_MemoryPath) Replace(t testing.TB, val *oc.Component_IntegratedCircuit_Memory) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit/memory in the given batch object.
func (n *Component_IntegratedCircuit_MemoryPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_Memory) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit/memory.
func (n *Component_IntegratedCircuit_MemoryPath) Update(t testing.TB, val *oc.Component_IntegratedCircuit_Memory) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit/memory in the given batch object.
func (n *Component_IntegratedCircuit_MemoryPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_Memory) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, true)
	if ok {
		return convertComponent_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/config/name with a ONCE subscription.
func (n *Component_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/config/name.
func (n *Component_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/config/name in the given batch object.
func (n *Component_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/config/name.
func (n *Component_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/config/name in the given batch object.
func (n *Component_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/config/name.
func (n *Component_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/config/name in the given batch object.
func (n *Component_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_NamePath extracts the value of the leaf Name from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PortPath) Lookup(t testing.TB) *oc.QualifiedComponent_Port {
	t.Helper()
	goStruct := &oc.Component_Port{}
	md, ok := oc.Lookup(t, n, "Component_Port", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Port{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PortPath) Get(t testing.TB) *oc.Component_Port {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port with a ONCE subscription.
func (n *Component_PortPathAny) Get(t testing.TB) []*oc.Component_Port {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port.
func (n *Component_PortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port in the given batch object.
func (n *Component_PortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port.
func (n *Component_PortPath) Replace(t testing.TB, val *oc.Component_Port) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port in the given batch object.
func (n *Component_PortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port.
func (n *Component_PortPath) Update(t testing.TB, val *oc.Component_Port) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port in the given batch object.
func (n *Component_PortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutModePath) Lookup(t testing.TB) *oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutModePath) Get(t testing.TB) *oc.Component_Port_BreakoutMode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutModePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port_BreakoutMode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
func (n *Component_Port_BreakoutModePathAny) Get(t testing.TB) []*oc.Component_Port_BreakoutMode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port_BreakoutMode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode.
func (n *Component_Port_BreakoutModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode in the given batch object.
func (n *Component_Port_BreakoutModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode.
func (n *Component_Port_BreakoutModePath) Replace(t testing.TB, val *oc.Component_Port_BreakoutMode) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode in the given batch object.
func (n *Component_Port_BreakoutModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode.
func (n *Component_Port_BreakoutModePath) Update(t testing.TB, val *oc.Component_Port_BreakoutMode) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode in the given batch object.
func (n *Component_Port_BreakoutModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_GroupPath) Lookup(t testing.TB) *oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_GroupPath) Get(t testing.TB) *oc.Component_Port_BreakoutMode_Group {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_GroupPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port_BreakoutMode_Group
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
func (n *Component_Port_BreakoutMode_GroupPathAny) Get(t testing.TB) []*oc.Component_Port_BreakoutMode_Group {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port_BreakoutMode_Group
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group.
func (n *Component_Port_BreakoutMode_GroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group in the given batch object.
func (n *Component_Port_BreakoutMode_GroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group.
func (n *Component_Port_BreakoutMode_GroupPath) Replace(t testing.TB, val *oc.Component_Port_BreakoutMode_Group) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group in the given batch object.
func (n *Component_Port_BreakoutMode_GroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode_Group) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group.
func (n *Component_Port_BreakoutMode_GroupPath) Update(t testing.TB, val *oc.Component_Port_BreakoutMode_Group) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group in the given batch object.
func (n *Component_Port_BreakoutMode_GroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode_Group) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, true)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Get(t testing.TB) oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Get(t testing.TB) []oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_ETHERNET_SPEED
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed in the given batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Replace(t testing.TB, val oc.E_IfEthernet_ETHERNET_SPEED) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed in the given batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_ETHERNET_SPEED) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Update(t testing.TB, val oc.E_IfEthernet_ETHERNET_SPEED) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed in the given batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_ETHERNET_SPEED) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath extracts the value of the leaf BreakoutSpeed from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_ETHERNET_SPEED.
func convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_ETHERNET_SPEED{
		Metadata: md,
	}
	val := parent.BreakoutSpeed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, true)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index in the given batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index in the given batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index in the given batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Port_BreakoutMode_Group_IndexPath extracts the value of the leaf Index from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
