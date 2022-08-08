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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_EgressSamplingRatePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_EgressSamplingRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_EgressSamplingRatePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_EgressSamplingRatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_EgressSamplingRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate with a ONCE subscription.
func (n *Sampling_Sflow_EgressSamplingRatePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate.
func (n *Sampling_Sflow_EgressSamplingRatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_EgressSamplingRatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate.
func (n *Sampling_Sflow_EgressSamplingRatePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_EgressSamplingRatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate.
func (n *Sampling_Sflow_EgressSamplingRatePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/egress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_EgressSamplingRatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_EgressSamplingRatePath extracts the value of the leaf EgressSamplingRate from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertSampling_Sflow_EgressSamplingRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.EgressSamplingRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/enabled with a ONCE subscription.
func (n *Sampling_Sflow_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/enabled.
func (n *Sampling_Sflow_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/enabled in the given batch object.
func (n *Sampling_Sflow_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/enabled.
func (n *Sampling_Sflow_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/enabled in the given batch object.
func (n *Sampling_Sflow_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/enabled.
func (n *Sampling_Sflow_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/enabled in the given batch object.
func (n *Sampling_Sflow_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_EnabledPath extracts the value of the leaf Enabled from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSampling_Sflow_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_IngressSamplingRatePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_IngressSamplingRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_IngressSamplingRatePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_IngressSamplingRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate with a ONCE subscription.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate.
func (n *Sampling_Sflow_IngressSamplingRatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_IngressSamplingRatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate.
func (n *Sampling_Sflow_IngressSamplingRatePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_IngressSamplingRatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate.
func (n *Sampling_Sflow_IngressSamplingRatePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/ingress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_IngressSamplingRatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_IngressSamplingRatePath extracts the value of the leaf IngressSamplingRate from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertSampling_Sflow_IngressSamplingRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.IngressSamplingRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_InterfacePath) Lookup(t testing.TB) *oc.QualifiedSampling_Sflow_Interface {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSampling_Sflow_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_InterfacePath) Get(t testing.TB) *oc.Sampling_Sflow_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedSampling_Sflow_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling_Sflow_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Interface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling_Sflow_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription.
func (n *Sampling_Sflow_InterfacePathAny) Get(t testing.TB) []*oc.Sampling_Sflow_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling_Sflow_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface.
func (n *Sampling_Sflow_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/interfaces/interface in the given batch object.
func (n *Sampling_Sflow_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface.
func (n *Sampling_Sflow_InterfacePath) Replace(t testing.TB, val *oc.Sampling_Sflow_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/interfaces/interface in the given batch object.
func (n *Sampling_Sflow_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling_Sflow_Interface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface.
func (n *Sampling_Sflow_InterfacePath) Update(t testing.TB, val *oc.Sampling_Sflow_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/interfaces/interface in the given batch object.
func (n *Sampling_Sflow_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Sampling_Sflow_Interface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Interface_EgressSamplingRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Interface_EgressSamplingRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate with a ONCE subscription.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/egress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Interface_EgressSamplingRatePath extracts the value of the leaf EgressSamplingRate from its parent oc.Sampling_Sflow_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertSampling_Sflow_Interface_EgressSamplingRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Interface) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.EgressSamplingRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
