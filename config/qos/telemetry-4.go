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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-packets in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath extracts the value of the leaf MaxQueueDepthPackets from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.MaxQueueDepthPackets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/max-queue-depth-percent in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath extracts the value of the leaf MaxQueueDepthPercent from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.MaxQueueDepthPercent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Lookup(t testing.TB) *oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Get(t testing.TB) oc.E_QosTypes_QueueBehavior {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Lookup(t testing.TB) []*oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_QosTypes_QueueBehavior
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Get(t testing.TB) []oc.E_QosTypes_QueueBehavior {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_QosTypes_QueueBehavior
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Replace(t testing.TB, val oc.E_QosTypes_QueueBehavior) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_QosTypes_QueueBehavior) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Update(t testing.TB, val oc.E_QosTypes_QueueBehavior) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/config/queuing-behavior in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_QosTypes_QueueBehavior) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath extracts the value of the leaf QueuingBehavior from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_QosTypes_QueueBehavior.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	qv := &oc.QualifiedE_QosTypes_QueueBehavior{
		Metadata: md,
	}
	val := parent.QueuingBehavior
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Output
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_Output
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_Output) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_Output) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_Output) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_Output) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/child-scheduler in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath extracts the value of the leaf ChildScheduler from its parent oc.Qos_SchedulerPolicy_Scheduler_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ChildScheduler
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-fwd-group in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath extracts the value of the leaf OutputFwdGroup from its parent oc.Qos_SchedulerPolicy_Scheduler_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OutputFwdGroup
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Lookup(t testing.TB) *oc.QualifiedE_Output_OutputType {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Get(t testing.TB) oc.E_Output_OutputType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Output_OutputType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Output_OutputType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Get(t testing.TB) []oc.E_Output_OutputType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Output_OutputType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Replace(t testing.TB, val oc.E_Output_OutputType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Output_OutputType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Update(t testing.TB, val oc.E_Output_OutputType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/config/output-type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Output_OutputType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath extracts the value of the leaf OutputType from its parent oc.Qos_SchedulerPolicy_Scheduler_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Output_OutputType.
func convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedE_Output_OutputType {
	t.Helper()
	qv := &oc.QualifiedE_Output_OutputType{
		Metadata: md,
	}
	val := parent.OutputType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Lookup(t testing.TB) *oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_PriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Get(t testing.TB) oc.E_Scheduler_Priority {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Scheduler_Priority
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Get(t testing.TB) []oc.E_Scheduler_Priority {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Scheduler_Priority
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Replace(t testing.TB, val oc.E_Scheduler_Priority) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Scheduler_Priority) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Update(t testing.TB, val oc.E_Scheduler_Priority) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/priority in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Scheduler_Priority) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_SchedulerPolicy_Scheduler_PriorityPath extracts the value of the leaf Priority from its parent oc.Qos_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Scheduler_Priority.
func convertQos_SchedulerPolicy_Scheduler_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler) *oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	qv := &oc.QualifiedE_Scheduler_Priority{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/sequence in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_SequencePath extracts the value of the leaf Sequence from its parent oc.Qos_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Sequence
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/bc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath extracts the value of the leaf Bc from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/be in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath extracts the value of the leaf Be from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Be
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath extracts the value of the leaf Cir from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath extracts the value of the leaf CirPct from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/cir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath extracts the value of the leaf CirPctRemaining from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath extracts the value of the leaf Drop from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath extracts the value of the leaf Pir from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Pir
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath extracts the value of the leaf PirPct from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PirPct
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/config/pir-pct-remaining in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath extracts the value of the leaf PirPctRemaining from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_PirPctRemainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PirPctRemaining
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) Replace(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) Update(t testing.TB, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/drop in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath extracts the value of the leaf Drop from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dot1p in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-dscp in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/violate-action/config/set-mpls-tc in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) Lookup(t testing.TB) *oc.QualifiedE_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, true, true)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) Get(t testing.TB) oc.E_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_QosTypes_QOS_SCHEDULER_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TypePathAny) Get(t testing.TB) []oc.E_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_QosTypes_QOS_SCHEDULER_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) Replace(t testing.TB, val oc.E_QosTypes_QOS_SCHEDULER_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_QosTypes_QOS_SCHEDULER_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) Update(t testing.TB, val oc.E_QosTypes_QOS_SCHEDULER_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/config/type in the given batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_QosTypes_QOS_SCHEDULER_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_SchedulerPolicy_Scheduler_TypePath extracts the value of the leaf Type from its parent oc.Qos_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedE_QosTypes_QOS_SCHEDULER_TYPE.
func convertQos_SchedulerPolicy_Scheduler_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler) *oc.QualifiedE_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_QosTypes_QOS_SCHEDULER_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
