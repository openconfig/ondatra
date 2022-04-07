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

// Lookup fetches the value at /openconfig-qos/qos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *QosPath) Lookup(t testing.TB) *oc.QualifiedQos {
	t.Helper()
	goStruct := &oc.Qos{}
	md, ok := oc.Lookup(t, n, "Qos", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *QosPath) Get(t testing.TB) *oc.Qos {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *QosPathAny) Lookup(t testing.TB) []*oc.QualifiedQos {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos with a ONCE subscription.
func (n *QosPathAny) Get(t testing.TB) []*oc.Qos {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos.
func (n *QosPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos in the given batch object.
func (n *QosPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos.
func (n *QosPath) Replace(t testing.TB, val *oc.Qos) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos in the given batch object.
func (n *QosPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos.
func (n *QosPath) Update(t testing.TB, val *oc.Qos) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos in the given batch object.
func (n *QosPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedQos_BufferAllocationProfile {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_BufferAllocationProfile{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfilePath) Get(t testing.TB) *oc.Qos_BufferAllocationProfile {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_BufferAllocationProfile {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_BufferAllocationProfile
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_BufferAllocationProfile{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription.
func (n *Qos_BufferAllocationProfilePathAny) Get(t testing.TB) []*oc.Qos_BufferAllocationProfile {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_BufferAllocationProfile
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile.
func (n *Qos_BufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile in the given batch object.
func (n *Qos_BufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile.
func (n *Qos_BufferAllocationProfilePath) Replace(t testing.TB, val *oc.Qos_BufferAllocationProfile) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile in the given batch object.
func (n *Qos_BufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_BufferAllocationProfile) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile.
func (n *Qos_BufferAllocationProfilePath) Update(t testing.TB, val *oc.Qos_BufferAllocationProfile) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile in the given batch object.
func (n *Qos_BufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_BufferAllocationProfile) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name.
func (n *Qos_BufferAllocationProfile_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name in the given batch object.
func (n *Qos_BufferAllocationProfile_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name.
func (n *Qos_BufferAllocationProfile_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name in the given batch object.
func (n *Qos_BufferAllocationProfile_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name.
func (n *Qos_BufferAllocationProfile_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/config/name in the given batch object.
func (n *Qos_BufferAllocationProfile_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_BufferAllocationProfile_NamePath extracts the value of the leaf Name from its parent oc.Qos_BufferAllocationProfile
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_BufferAllocationProfile_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_BufferAllocationProfile_Queue {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_BufferAllocationProfile_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_QueuePath) Get(t testing.TB) *oc.Qos_BufferAllocationProfile_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_BufferAllocationProfile_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_BufferAllocationProfile_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_BufferAllocationProfile_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Get(t testing.TB) []*oc.Qos_BufferAllocationProfile_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_BufferAllocationProfile_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue.
func (n *Qos_BufferAllocationProfile_QueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue in the given batch object.
func (n *Qos_BufferAllocationProfile_QueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue.
func (n *Qos_BufferAllocationProfile_QueuePath) Replace(t testing.TB, val *oc.Qos_BufferAllocationProfile_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue in the given batch object.
func (n *Qos_BufferAllocationProfile_QueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_BufferAllocationProfile_Queue) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue.
func (n *Qos_BufferAllocationProfile_QueuePath) Update(t testing.TB, val *oc.Qos_BufferAllocationProfile_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue in the given batch object.
func (n *Qos_BufferAllocationProfile_QueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_BufferAllocationProfile_Queue) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dedicated-buffer in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath extracts the value of the leaf DedicatedBuffer from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DedicatedBuffer
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Replace(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Update(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/dynamic-limit-scaling-factor in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath extracts the value of the leaf DynamicLimitScalingFactor from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.DynamicLimitScalingFactor
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/name in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_BufferAllocationProfile_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_BufferAllocationProfile_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Lookup(t testing.TB) *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Get(t testing.TB) oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Get(t testing.TB) []oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Replace(t testing.TB, val oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Update(t testing.TB, val oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/shared-buffer-limit-type in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath extracts the value of the leaf SharedBufferLimitType from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE.
func convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE{
		Metadata: md,
	}
	val := parent.SharedBufferLimitType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/static-shared-buffer-limit in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath extracts the value of the leaf StaticSharedBufferLimit from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.StaticSharedBufferLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, true)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/config/use-shared-buffer in the given batch object.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath extracts the value of the leaf UseSharedBuffer from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.UseSharedBuffer
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ClassifierPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ClassifierPath) Get(t testing.TB) *oc.Qos_Classifier {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ClassifierPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription.
func (n *Qos_ClassifierPathAny) Get(t testing.TB) []*oc.Qos_Classifier {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier.
func (n *Qos_ClassifierPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier in the given batch object.
func (n *Qos_ClassifierPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier.
func (n *Qos_ClassifierPath) Replace(t testing.TB, val *oc.Qos_Classifier) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier in the given batch object.
func (n *Qos_ClassifierPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier.
func (n *Qos_ClassifierPath) Update(t testing.TB, val *oc.Qos_Classifier) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier in the given batch object.
func (n *Qos_ClassifierPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, true, true)
	if ok {
		return convertQos_Classifier_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/config/name with a ONCE subscription.
func (n *Qos_Classifier_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/config/name.
func (n *Qos_Classifier_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Classifier_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/config/name.
func (n *Qos_Classifier_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Classifier_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/config/name.
func (n *Qos_Classifier_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Classifier_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_NamePath extracts the value of the leaf Name from its parent oc.Qos_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_TermPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_TermPath) Get(t testing.TB) *oc.Qos_Classifier_Term {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_TermPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
func (n *Qos_Classifier_TermPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term.
func (n *Qos_Classifier_TermPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term in the given batch object.
func (n *Qos_Classifier_TermPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term.
func (n *Qos_Classifier_TermPath) Replace(t testing.TB, val *oc.Qos_Classifier_Term) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term in the given batch object.
func (n *Qos_Classifier_TermPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term.
func (n *Qos_Classifier_TermPath) Update(t testing.TB, val *oc.Qos_Classifier_Term) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term in the given batch object.
func (n *Qos_Classifier_TermPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_ActionsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_ActionsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
func (n *Qos_Classifier_Term_ActionsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions.
func (n *Qos_Classifier_Term_ActionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions in the given batch object.
func (n *Qos_Classifier_Term_ActionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions.
func (n *Qos_Classifier_Term_ActionsPath) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Actions) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions in the given batch object.
func (n *Qos_Classifier_Term_ActionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Actions) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions.
func (n *Qos_Classifier_Term_ActionsPath) Update(t testing.TB, val *oc.Qos_Classifier_Term_Actions) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions in the given batch object.
func (n *Qos_Classifier_Term_ActionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Actions) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Actions_Remark
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Actions_Remark
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark in the given batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Actions_Remark) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark in the given batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Actions_Remark) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Update(t testing.TB, val *oc.Qos_Classifier_Term_Actions_Remark) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark in the given batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Actions_Remark) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dot1p in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Actions_Remark_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-dscp in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Actions_Remark_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/config/set-mpls-tc in the given batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group in the given batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group in the given batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/config/target-group in the given batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Actions_TargetGroupPath extracts the value of the leaf TargetGroup from its parent oc.Qos_Classifier_Term_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Actions_TargetGroupPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TargetGroup
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_ConditionsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_ConditionsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_ConditionsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
func (n *Qos_Classifier_Term_ConditionsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions.
func (n *Qos_Classifier_Term_ConditionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions in the given batch object.
func (n *Qos_Classifier_Term_ConditionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions.
func (n *Qos_Classifier_Term_ConditionsPath) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Conditions) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions in the given batch object.
func (n *Qos_Classifier_Term_ConditionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions.
func (n *Qos_Classifier_Term_ConditionsPath) Update(t testing.TB, val *oc.Qos_Classifier_Term_Conditions) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions in the given batch object.
func (n *Qos_Classifier_Term_ConditionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Ipv4) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Ipv4) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Update(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Ipv4) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Ipv4) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/destination-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_DscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_DscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv4_DscpPath extracts the value of the leaf Dscp from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Ipv4_DscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Dscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) Lookup(t testing.TB) *oc.QualifiedUint8Slice {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_DscpSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) Get(t testing.TB) []uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_DscpSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPathAny) Get(t testing.TB) [][]uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) Replace(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) Update(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/dscp-set in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DscpSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Ipv4_DscpSetPath extracts the value of the leaf DscpSet from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8Slice.
func convertQos_Classifier_Term_Conditions_Ipv4_DscpSetPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedUint8Slice {
	t.Helper()
	qv := &oc.QualifiedUint8Slice{
		Metadata: md,
	}
	val := parent.DscpSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_HopLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_HopLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/hop-limit in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_HopLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv4_HopLimitPath extracts the value of the leaf HopLimit from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Ipv4_HopLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.HopLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_ProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_ProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/protocol in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_ProtocolPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Ipv4_ProtocolPath extracts the value of the leaf Protocol from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union.
func convertQos_Classifier_Term_Conditions_Ipv4_ProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union{
		Metadata: md,
	}
	val := parent.Protocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/config/source-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv4_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_Ipv4_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6PathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Ipv6) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) Update(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Ipv6) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_Ipv6_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/destination-flow-label in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath extracts the value of the leaf DestinationFlowLabel from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Classifier_Term_Conditions_Ipv6_DestinationFlowLabelPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.DestinationFlowLabel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_DscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_DscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_DscpPath extracts the value of the leaf Dscp from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Ipv6_DscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Dscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) Lookup(t testing.TB) *oc.QualifiedUint8Slice {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_DscpSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) Get(t testing.TB) []uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_DscpSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPathAny) Get(t testing.TB) [][]uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) Replace(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) Update(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/dscp-set in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_DscpSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_DscpSetPath extracts the value of the leaf DscpSet from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8Slice.
func convertQos_Classifier_Term_Conditions_Ipv6_DscpSetPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedUint8Slice {
	t.Helper()
	qv := &oc.QualifiedUint8Slice{
		Metadata: md,
	}
	val := parent.DscpSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_HopLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_HopLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/hop-limit in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_HopLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_HopLimitPath extracts the value of the leaf HopLimit from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Ipv6_HopLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.HopLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/protocol in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath extracts the value of the leaf Protocol from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union.
func convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union{
		Metadata: md,
	}
	val := parent.Protocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
