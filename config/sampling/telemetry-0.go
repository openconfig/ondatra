package sampling

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

// Lookup fetches the value at /openconfig-sampling/sampling with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SamplingPath) Lookup(t testing.TB) *oc.QualifiedSampling {
	t.Helper()
	goStruct := &oc.Sampling{}
	md, ok := oc.Lookup(t, n, "Sampling", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSampling{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SamplingPath) Get(t testing.TB) *oc.Sampling {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SamplingPathAny) Lookup(t testing.TB) []*oc.QualifiedSampling {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling with a ONCE subscription.
func (n *SamplingPathAny) Get(t testing.TB) []*oc.Sampling {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling.
func (n *SamplingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling in the given batch object.
func (n *SamplingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling.
func (n *SamplingPath) Replace(t testing.TB, val *oc.Sampling) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling in the given batch object.
func (n *SamplingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-sampling/sampling.
func (n *SamplingPath) Update(t testing.TB, val *oc.Sampling) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling in the given batch object.
func (n *SamplingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_SflowPath) Lookup(t testing.TB) *oc.QualifiedSampling_Sflow {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSampling_Sflow{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_SflowPath) Get(t testing.TB) *oc.Sampling_Sflow {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_SflowPathAny) Lookup(t testing.TB) []*oc.QualifiedSampling_Sflow {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling_Sflow
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling_Sflow{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow with a ONCE subscription.
func (n *Sampling_SflowPathAny) Get(t testing.TB) []*oc.Sampling_Sflow {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling_Sflow
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow.
func (n *Sampling_SflowPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow in the given batch object.
func (n *Sampling_SflowPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow.
func (n *Sampling_SflowPath) Replace(t testing.TB, val *oc.Sampling_Sflow) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow in the given batch object.
func (n *Sampling_SflowPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling_Sflow) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow.
func (n *Sampling_SflowPath) Update(t testing.TB, val *oc.Sampling_Sflow) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow in the given batch object.
func (n *Sampling_SflowPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling_Sflow) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_AgentIdIpv4Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_AgentIdIpv4Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_AgentIdIpv4Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_AgentIdIpv4Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 with a ONCE subscription.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4.
func (n *Sampling_Sflow_AgentIdIpv4Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 in the given batch object.
func (n *Sampling_Sflow_AgentIdIpv4Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4.
func (n *Sampling_Sflow_AgentIdIpv4Path) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 in the given batch object.
func (n *Sampling_Sflow_AgentIdIpv4Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4.
func (n *Sampling_Sflow_AgentIdIpv4Path) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/agent-id-ipv4 in the given batch object.
func (n *Sampling_Sflow_AgentIdIpv4Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_AgentIdIpv4Path extracts the value of the leaf AgentIdIpv4 from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_AgentIdIpv4Path(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AgentIdIpv4
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_AgentIdIpv6Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_AgentIdIpv6Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_AgentIdIpv6Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_AgentIdIpv6Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 with a ONCE subscription.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6.
func (n *Sampling_Sflow_AgentIdIpv6Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 in the given batch object.
func (n *Sampling_Sflow_AgentIdIpv6Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6.
func (n *Sampling_Sflow_AgentIdIpv6Path) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 in the given batch object.
func (n *Sampling_Sflow_AgentIdIpv6Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6.
func (n *Sampling_Sflow_AgentIdIpv6Path) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/agent-id-ipv6 in the given batch object.
func (n *Sampling_Sflow_AgentIdIpv6Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_AgentIdIpv6Path extracts the value of the leaf AgentIdIpv6 from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_AgentIdIpv6Path(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AgentIdIpv6
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_CollectorPath) Lookup(t testing.TB) *oc.QualifiedSampling_Sflow_Collector {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSampling_Sflow_Collector{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_CollectorPath) Get(t testing.TB) *oc.Sampling_Sflow_Collector {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_CollectorPathAny) Lookup(t testing.TB) []*oc.QualifiedSampling_Sflow_Collector {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling_Sflow_Collector
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling_Sflow_Collector{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription.
func (n *Sampling_Sflow_CollectorPathAny) Get(t testing.TB) []*oc.Sampling_Sflow_Collector {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling_Sflow_Collector
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/collectors/collector.
func (n *Sampling_Sflow_CollectorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/collectors/collector in the given batch object.
func (n *Sampling_Sflow_CollectorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/collectors/collector.
func (n *Sampling_Sflow_CollectorPath) Replace(t testing.TB, val *oc.Sampling_Sflow_Collector) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/collectors/collector in the given batch object.
func (n *Sampling_Sflow_CollectorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling_Sflow_Collector) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/collectors/collector.
func (n *Sampling_Sflow_CollectorPath) Update(t testing.TB, val *oc.Sampling_Sflow_Collector) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/collectors/collector in the given batch object.
func (n *Sampling_Sflow_CollectorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling_Sflow_Collector) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
