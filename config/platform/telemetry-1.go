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
// failing the test fatally if no value is present at the path.
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
// failing the test fatally if no value is present at the path.
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

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/config/node-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_NodeIdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit", goStruct, true, true)
	if ok {
		return convertComponent_IntegratedCircuit_NodeIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/config/node-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_NodeIdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/config/node-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_NodeIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/config/node-id with a ONCE subscription.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit/config/node-id.
func (n *Component_IntegratedCircuit_NodeIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit/config/node-id in the given batch object.
func (n *Component_IntegratedCircuit_NodeIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit/config/node-id.
func (n *Component_IntegratedCircuit_NodeIdPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit/config/node-id in the given batch object.
func (n *Component_IntegratedCircuit_NodeIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit/config/node-id.
func (n *Component_IntegratedCircuit_NodeIdPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit/config/node-id in the given batch object.
func (n *Component_IntegratedCircuit_NodeIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_IntegratedCircuit_NodeIdPath extracts the value of the leaf NodeId from its parent oc.Component_IntegratedCircuit
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_NodeIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.NodeId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_UtilizationPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_UtilizationPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Utilization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Utilization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Utilization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Utilization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit/utilization.
func (n *Component_IntegratedCircuit_UtilizationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit/utilization in the given batch object.
func (n *Component_IntegratedCircuit_UtilizationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit/utilization.
func (n *Component_IntegratedCircuit_UtilizationPath) Replace(t testing.TB, val *oc.Component_IntegratedCircuit_Utilization) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit/utilization in the given batch object.
func (n *Component_IntegratedCircuit_UtilizationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_Utilization) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit/utilization.
func (n *Component_IntegratedCircuit_UtilizationPath) Update(t testing.TB, val *oc.Component_IntegratedCircuit_Utilization) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit/utilization in the given batch object.
func (n *Component_IntegratedCircuit_UtilizationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_Utilization) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Utilization_Resource
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource in the given batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Replace(t testing.TB, val *oc.Component_IntegratedCircuit_Utilization_Resource) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource in the given batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_Utilization_Resource) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Update(t testing.TB, val *oc.Component_IntegratedCircuit_Utilization_Resource) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource in the given batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_IntegratedCircuit_Utilization_Resource) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, true)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name in the given batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name in the given batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/config/name in the given batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_NamePath extracts the value of the leaf Name from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedString {
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
