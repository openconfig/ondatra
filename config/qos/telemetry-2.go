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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_Queue_QueueManagementProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_QueueManagementProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile in the given batch object.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile in the given batch object.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/queue-management-profile in the given batch object.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_Queue_QueueManagementProfilePath extracts the value of the leaf QueueManagementProfile from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Queue_QueueManagementProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.QueueManagementProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Get(t testing.TB) *oc.Qos_Interface_Input_SchedulerPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_SchedulerPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_SchedulerPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_SchedulerPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy in the given batch object.
func (n *Qos_Interface_Input_SchedulerPolicyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Replace(t testing.TB, val *oc.Qos_Interface_Input_SchedulerPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy in the given batch object.
func (n *Qos_Interface_Input_SchedulerPolicyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_SchedulerPolicy) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Update(t testing.TB, val *oc.Qos_Interface_Input_SchedulerPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy in the given batch object.
func (n *Qos_Interface_Input_SchedulerPolicyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_SchedulerPolicy) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name in the given batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name in the given batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/config/name in the given batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_SchedulerPolicy_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_SchedulerPolicy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_SchedulerPolicy_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/config/unicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_UnicastBufferAllocationProfilePath extracts the value of the leaf UnicastBufferAllocationProfile from its parent oc.Qos_Interface_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.UnicastBufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterfacePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_VoqInterface {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_VoqInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterfacePath) Get(t testing.TB) *oc.Qos_Interface_Input_VoqInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_VoqInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_VoqInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_VoqInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_VoqInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_VoqInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface.
func (n *Qos_Interface_Input_VoqInterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface in the given batch object.
func (n *Qos_Interface_Input_VoqInterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface.
func (n *Qos_Interface_Input_VoqInterfacePath) Replace(t testing.TB, val *oc.Qos_Interface_Input_VoqInterface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface in the given batch object.
func (n *Qos_Interface_Input_VoqInterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_VoqInterface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface.
func (n *Qos_Interface_Input_VoqInterfacePath) Update(t testing.TB, val *oc.Qos_Interface_Input_VoqInterface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface in the given batch object.
func (n *Qos_Interface_Input_VoqInterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_VoqInterface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_VoqInterface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/config/name in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_VoqInterface_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_VoqInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_VoqInterface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Get(t testing.TB) *oc.Qos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_VoqInterface_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_VoqInterface_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Replace(t testing.TB, val *oc.Qos_Interface_Input_VoqInterface_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_VoqInterface_Queue) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Update(t testing.TB, val *oc.Qos_Interface_Input_VoqInterface_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_VoqInterface_Queue) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_VoqInterface_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_VoqInterface_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/config/interface-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface", goStruct, true, true)
	if ok {
		return convertQos_Interface_InterfaceIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/config/interface-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/config/interface-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_InterfaceIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/config/interface-id with a ONCE subscription.
func (n *Qos_Interface_InterfaceIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/config/interface-id.
func (n *Qos_Interface_InterfaceIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/config/interface-id in the given batch object.
func (n *Qos_Interface_InterfaceIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/config/interface-id.
func (n *Qos_Interface_InterfaceIdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/config/interface-id in the given batch object.
func (n *Qos_Interface_InterfaceIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/config/interface-id.
func (n *Qos_Interface_InterfaceIdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/config/interface-id in the given batch object.
func (n *Qos_Interface_InterfaceIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_InterfaceIdPath extracts the value of the leaf InterfaceId from its parent oc.Qos_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_InterfaceIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.InterfaceId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_InterfaceRef {
	t.Helper()
	goStruct := &oc.Qos_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceRefPath) Get(t testing.TB) *oc.Qos_Interface_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription.
func (n *Qos_Interface_InterfaceRefPathAny) Get(t testing.TB) []*oc.Qos_Interface_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref.
func (n *Qos_Interface_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/interface-ref in the given batch object.
func (n *Qos_Interface_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref.
func (n *Qos_Interface_InterfaceRefPath) Replace(t testing.TB, val *oc.Qos_Interface_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/interface-ref in the given batch object.
func (n *Qos_Interface_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref.
func (n *Qos_Interface_InterfaceRefPath) Update(t testing.TB, val *oc.Qos_Interface_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/interface-ref in the given batch object.
func (n *Qos_Interface_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_InterfaceRef", goStruct, true, true)
	if ok {
		return convertQos_Interface_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface with a ONCE subscription.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface in the given batch object.
func (n *Qos_Interface_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface in the given batch object.
func (n *Qos_Interface_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/interface-ref/config/interface in the given batch object.
func (n *Qos_Interface_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Qos_Interface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_InterfaceRef) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_InterfaceRef", goStruct, true, true)
	if ok {
		return convertQos_Interface_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface in the given batch object.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface in the given batch object.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/interface-ref/config/subinterface in the given batch object.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Qos_Interface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Interface_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_InterfaceRef) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_OutputPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Output{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_OutputPath) Get(t testing.TB) *oc.Qos_Interface_Output {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_OutputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription.
func (n *Qos_Interface_OutputPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output.
func (n *Qos_Interface_OutputPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output in the given batch object.
func (n *Qos_Interface_OutputPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output.
func (n *Qos_Interface_OutputPath) Replace(t testing.TB, val *oc.Qos_Interface_Output) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output in the given batch object.
func (n *Qos_Interface_OutputPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output.
func (n *Qos_Interface_OutputPath) Update(t testing.TB, val *oc.Qos_Interface_Output) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output in the given batch object.
func (n *Qos_Interface_OutputPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_BufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_BufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/config/buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_BufferAllocationProfilePath extracts the value of the leaf BufferAllocationProfile from its parent oc.Qos_Interface_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_BufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.BufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_ClassifierPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_Classifier {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Classifier", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_ClassifierPath) Get(t testing.TB) *oc.Qos_Interface_Output_Classifier {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_ClassifierPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_Classifier {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_Classifier
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Classifier", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription.
func (n *Qos_Interface_Output_ClassifierPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_Classifier {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_Classifier
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier.
func (n *Qos_Interface_Output_ClassifierPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier in the given batch object.
func (n *Qos_Interface_Output_ClassifierPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier.
func (n *Qos_Interface_Output_ClassifierPath) Replace(t testing.TB, val *oc.Qos_Interface_Output_Classifier) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier in the given batch object.
func (n *Qos_Interface_Output_ClassifierPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_Classifier) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier.
func (n *Qos_Interface_Output_ClassifierPath) Update(t testing.TB, val *oc.Qos_Interface_Output_Classifier) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier in the given batch object.
func (n *Qos_Interface_Output_ClassifierPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_Classifier) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Classifier_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Classifier", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_Classifier_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Classifier_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Classifier_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Classifier", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Classifier_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name with a ONCE subscription.
func (n *Qos_Interface_Output_Classifier_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name.
func (n *Qos_Interface_Output_Classifier_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Interface_Output_Classifier_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name.
func (n *Qos_Interface_Output_Classifier_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Interface_Output_Classifier_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name.
func (n *Qos_Interface_Output_Classifier_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Interface_Output_Classifier_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_Classifier_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Output_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_Classifier_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Classifier) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Classifier_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Classifier", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_Classifier_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Classifier_TypePath) Get(t testing.TB) oc.E_Input_Classifier_Type {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Input_Classifier_Type
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Classifier", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Classifier_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type with a ONCE subscription.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Get(t testing.TB) []oc.E_Input_Classifier_Type {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Input_Classifier_Type
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type.
func (n *Qos_Interface_Output_Classifier_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Interface_Output_Classifier_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type.
func (n *Qos_Interface_Output_Classifier_TypePath) Replace(t testing.TB, val oc.E_Input_Classifier_Type) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Interface_Output_Classifier_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Input_Classifier_Type) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type.
func (n *Qos_Interface_Output_Classifier_TypePath) Update(t testing.TB, val oc.E_Input_Classifier_Type) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Interface_Output_Classifier_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Input_Classifier_Type) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Interface_Output_Classifier_TypePath extracts the value of the leaf Type from its parent oc.Qos_Interface_Output_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Input_Classifier_Type.
func convertQos_Interface_Output_Classifier_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Classifier) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	qv := &oc.QualifiedE_Input_Classifier_Type{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_InterfaceRef {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_InterfaceRefPath) Get(t testing.TB) *oc.Qos_Interface_Output_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref.
func (n *Qos_Interface_Output_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref in the given batch object.
func (n *Qos_Interface_Output_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref.
func (n *Qos_Interface_Output_InterfaceRefPath) Replace(t testing.TB, val *oc.Qos_Interface_Output_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref in the given batch object.
func (n *Qos_Interface_Output_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref.
func (n *Qos_Interface_Output_InterfaceRefPath) Update(t testing.TB, val *oc.Qos_Interface_Output_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref in the given batch object.
func (n *Qos_Interface_Output_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_InterfaceRef", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface with a ONCE subscription.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface in the given batch object.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface in the given batch object.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/interface in the given batch object.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Qos_Interface_Output_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_InterfaceRef) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_InterfaceRef", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface with a ONCE subscription.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface in the given batch object.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface in the given batch object.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/interface-ref/config/subinterface in the given batch object.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Qos_Interface_Output_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_InterfaceRef) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/config/multicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_MulticastBufferAllocationProfilePath extracts the value of the leaf MulticastBufferAllocationProfile from its parent oc.Qos_Interface_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MulticastBufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_Queue {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_QueuePath) Get(t testing.TB) *oc.Qos_Interface_Output_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription.
func (n *Qos_Interface_Output_QueuePathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue.
func (n *Qos_Interface_Output_QueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue in the given batch object.
func (n *Qos_Interface_Output_QueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue.
func (n *Qos_Interface_Output_QueuePath) Replace(t testing.TB, val *oc.Qos_Interface_Output_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue in the given batch object.
func (n *Qos_Interface_Output_QueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_Queue) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue.
func (n *Qos_Interface_Output_QueuePath) Update(t testing.TB, val *oc.Qos_Interface_Output_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue in the given batch object.
func (n *Qos_Interface_Output_QueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_Queue) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name.
func (n *Qos_Interface_Output_Queue_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Output_Queue_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name.
func (n *Qos_Interface_Output_Queue_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Output_Queue_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name.
func (n *Qos_Interface_Output_Queue_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Output_Queue_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_Queue_QueueManagementProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_QueueManagementProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile in the given batch object.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile in the given batch object.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/queues/queue/config/queue-management-profile in the given batch object.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_Queue_QueueManagementProfilePath extracts the value of the leaf QueueManagementProfile from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_Queue_QueueManagementProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.QueueManagementProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_SchedulerPolicyPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_SchedulerPolicy {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_SchedulerPolicy", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_SchedulerPolicyPath) Get(t testing.TB) *oc.Qos_Interface_Output_SchedulerPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_SchedulerPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_SchedulerPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_SchedulerPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy with a ONCE subscription.
func (n *Qos_Interface_Output_SchedulerPolicyPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_SchedulerPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_SchedulerPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy.
func (n *Qos_Interface_Output_SchedulerPolicyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy in the given batch object.
func (n *Qos_Interface_Output_SchedulerPolicyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy.
func (n *Qos_Interface_Output_SchedulerPolicyPath) Replace(t testing.TB, val *oc.Qos_Interface_Output_SchedulerPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy in the given batch object.
func (n *Qos_Interface_Output_SchedulerPolicyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_SchedulerPolicy) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy.
func (n *Qos_Interface_Output_SchedulerPolicyPath) Update(t testing.TB, val *oc.Qos_Interface_Output_SchedulerPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy in the given batch object.
func (n *Qos_Interface_Output_SchedulerPolicyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Output_SchedulerPolicy) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_SchedulerPolicy", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_SchedulerPolicy_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_SchedulerPolicy_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name with a ONCE subscription.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name in the given batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name in the given batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/config/name in the given batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_SchedulerPolicy_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Output_SchedulerPolicy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_SchedulerPolicy_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_SchedulerPolicy) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, true, true)
	if ok {
		return convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/output/config/unicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Output_UnicastBufferAllocationProfilePath extracts the value of the leaf UnicastBufferAllocationProfile from its parent oc.Qos_Interface_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.UnicastBufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfilePath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfilePath) Get(t testing.TB) *oc.Qos_QueueManagementProfile {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription.
func (n *Qos_QueueManagementProfilePathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile.
func (n *Qos_QueueManagementProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile in the given batch object.
func (n *Qos_QueueManagementProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile.
func (n *Qos_QueueManagementProfilePath) Replace(t testing.TB, val *oc.Qos_QueueManagementProfile) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile in the given batch object.
func (n *Qos_QueueManagementProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile.
func (n *Qos_QueueManagementProfilePath) Update(t testing.TB, val *oc.Qos_QueueManagementProfile) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile in the given batch object.
func (n *Qos_QueueManagementProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name with a ONCE subscription.
func (n *Qos_QueueManagementProfile_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name.
func (n *Qos_QueueManagementProfile_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name in the given batch object.
func (n *Qos_QueueManagementProfile_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name.
func (n *Qos_QueueManagementProfile_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name in the given batch object.
func (n *Qos_QueueManagementProfile_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name.
func (n *Qos_QueueManagementProfile_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/config/name in the given batch object.
func (n *Qos_QueueManagementProfile_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_NamePath extracts the value of the leaf Name from its parent oc.Qos_QueueManagementProfile
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_QueueManagementProfile_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_RedPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Red {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Red{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_RedPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Red {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_RedPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Red {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Red
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Red{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription.
func (n *Qos_QueueManagementProfile_RedPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Red {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Red
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red.
func (n *Qos_QueueManagementProfile_RedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red in the given batch object.
func (n *Qos_QueueManagementProfile_RedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red.
func (n *Qos_QueueManagementProfile_RedPath) Replace(t testing.TB, val *oc.Qos_QueueManagementProfile_Red) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red in the given batch object.
func (n *Qos_QueueManagementProfile_RedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Red) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red.
func (n *Qos_QueueManagementProfile_RedPath) Update(t testing.TB, val *oc.Qos_QueueManagementProfile_Red) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red in the given batch object.
func (n *Qos_QueueManagementProfile_RedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Red) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Red_Uniform
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Red_Uniform
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform in the given batch object.
func (n *Qos_QueueManagementProfile_Red_UniformPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Replace(t testing.TB, val *oc.Qos_QueueManagementProfile_Red_Uniform) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform in the given batch object.
func (n *Qos_QueueManagementProfile_Red_UniformPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Red_Uniform) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Update(t testing.TB, val *oc.Qos_QueueManagementProfile_Red_Uniform) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform in the given batch object.
func (n *Qos_QueueManagementProfile_Red_UniformPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_QueueManagementProfile_Red_Uniform) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_DropPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDrop())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/drop in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Red_Uniform_DropPath extracts the value of the leaf Drop from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Red_Uniform_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableEcn())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/enable-ecn in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath extracts the value of the leaf EnableEcn from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, true)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/config/max-threshold in the given batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath extracts the value of the leaf MaxThreshold from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedUint64 {
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
