package qos

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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/min-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath extracts the value of the leaf MinThreshold from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_WredPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Wred {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Wred{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_WredPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Wred {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_WredPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Wred {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Wred
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Wred{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription.
func (n *Qos_QueueManagementProfile_WredPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Wred {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Wred
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred.
func (n *Qos_QueueManagementProfile_WredPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred in the given batch object.
func (n *Qos_QueueManagementProfile_WredPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred.
func (n *Qos_QueueManagementProfile_WredPath) Replace(t testing.TB, val *oc.Qos_QueueManagementProfile_Wred) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred in the given batch object.
func (n *Qos_QueueManagementProfile_WredPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Wred) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred.
func (n *Qos_QueueManagementProfile_WredPath) Update(t testing.TB, val *oc.Qos_QueueManagementProfile_Wred) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred in the given batch object.
func (n *Qos_QueueManagementProfile_WredPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Wred) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Wred_Uniform
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Wred_Uniform
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Replace(t testing.TB, val *oc.Qos_QueueManagementProfile_Wred_Uniform) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Wred_Uniform) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Update(t testing.TB, val *oc.Qos_QueueManagementProfile_Wred_Uniform) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Wred_Uniform) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDrop())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/drop in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Wred_Uniform_DropPath extracts the value of the leaf Drop from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Drop
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableEcn())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/enable-ecn in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath extracts the value of the leaf EnableEcn from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableEcn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-drop-probability-percent in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath extracts the value of the leaf MaxDropProbabilityPercent from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.MaxDropProbabilityPercent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/max-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath extracts the value of the leaf MaxThreshold from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/min-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath extracts the value of the leaf MinThreshold from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/config/weight in the given batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Wred_Uniform_WeightPath extracts the value of the leaf Weight from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Weight
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Queue {
	t.Helper()
	goStruct := &oc.Qos_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Queue", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queues/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueuePath) Get(t testing.TB) *oc.Qos_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Queue", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queues/queue with a ONCE subscription.
func (n *Qos_QueuePathAny) Get(t testing.TB) []*oc.Qos_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queues/queue.
func (n *Qos_QueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queues/queue in the given batch object.
func (n *Qos_QueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queues/queue.
func (n *Qos_QueuePath) Replace(t testing.TB, val *oc.Qos_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queues/queue in the given batch object.
func (n *Qos_QueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Queue) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/queues/queue.
func (n *Qos_QueuePath) Update(t testing.TB, val *oc.Qos_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queues/queue in the given batch object.
func (n *Qos_QueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Queue) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/queues/queue/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Queue", goStruct, true, true)
	if ok {
		return convertQos_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queues/queue/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queues/queue/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queues/queue/config/name with a ONCE subscription.
func (n *Qos_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queues/queue/config/name.
func (n *Qos_Queue_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queues/queue/config/name in the given batch object.
func (n *Qos_Queue_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queues/queue/config/name.
func (n *Qos_Queue_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queues/queue/config/name in the given batch object.
func (n *Qos_Queue_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queues/queue/config/name.
func (n *Qos_Queue_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queues/queue/config/name in the given batch object.
func (n *Qos_Queue_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicyPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicyPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription.
func (n *Qos_SchedulerPolicyPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy.
func (n *Qos_SchedulerPolicyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy in the given batch object.
func (n *Qos_SchedulerPolicyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy.
func (n *Qos_SchedulerPolicyPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy in the given batch object.
func (n *Qos_SchedulerPolicyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy.
func (n *Qos_SchedulerPolicyPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy in the given batch object.
func (n *Qos_SchedulerPolicyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name with a ONCE subscription.
func (n *Qos_SchedulerPolicy_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name.
func (n *Qos_SchedulerPolicy_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name in the given batch object.
func (n *Qos_SchedulerPolicy_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name.
func (n *Qos_SchedulerPolicy_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name in the given batch object.
func (n *Qos_SchedulerPolicy_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name.
func (n *Qos_SchedulerPolicy_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/config/name in the given batch object.
func (n *Qos_SchedulerPolicy_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_NamePath extracts the value of the leaf Name from its parent oc.Qos_SchedulerPolicy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_SchedulerPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_SchedulerPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler.
func (n *Qos_SchedulerPolicy_SchedulerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler in the given batch object.
func (n *Qos_SchedulerPolicy_SchedulerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler.
func (n *Qos_SchedulerPolicy_SchedulerPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler in the given batch object.
func (n *Qos_SchedulerPolicy_SchedulerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler.
func (n *Qos_SchedulerPolicy_SchedulerPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler in the given batch object.
func (n *Qos_SchedulerPolicy_SchedulerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Input
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_Input
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_Input) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_Input) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_Input) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_Input) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/id in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_Input_IdPath extracts the value of the leaf Id from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Lookup(t testing.TB) *oc.QualifiedE_Input_InputType {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Get(t testing.TB) oc.E_Input_InputType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Input_InputType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Input_InputType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Get(t testing.TB) []oc.E_Input_InputType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Input_InputType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Replace(t testing.TB, val oc.E_Input_InputType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Input_InputType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Update(t testing.TB, val oc.E_Input_InputType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/input-type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Input_InputType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath extracts the value of the leaf InputType from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Input_InputType.
func convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedE_Input_InputType {
	t.Helper()
	qv := &oc.QualifiedE_Input_InputType{
		Metadata: md,
	}
	val := parent.InputType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/queue in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_Input_QueuePath extracts the value of the leaf Queue from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Queue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/config/weight in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_Input_WeightPath extracts the value of the leaf Weight from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Weight
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/bc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath extracts the value of the leaf Bc from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Bc
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath extracts the value of the leaf Cir from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Cir
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath extracts the value of the leaf CirPct from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CirPct
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/cir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath extracts the value of the leaf CirPctRemaining from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CirPctRemaining
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDot1P
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetMplsTc
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath extracts the value of the leaf Drop from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Drop
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDot1P
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetMplsTc
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-bytes in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath extracts the value of the leaf MaxQueueDepthBytes from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.MaxQueueDepthBytes
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
