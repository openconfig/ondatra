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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Interface_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Interface_EnabledPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Interface_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Interface_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Interface_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled with a ONCE subscription.
func (n *Sampling_Sflow_Interface_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled.
func (n *Sampling_Sflow_Interface_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled in the given batch object.
func (n *Sampling_Sflow_Interface_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled.
func (n *Sampling_Sflow_Interface_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled in the given batch object.
func (n *Sampling_Sflow_Interface_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled.
func (n *Sampling_Sflow_Interface_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/enabled in the given batch object.
func (n *Sampling_Sflow_Interface_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Interface_EnabledPath extracts the value of the leaf Enabled from its parent oc.Sampling_Sflow_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSampling_Sflow_Interface_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Interface) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Interface_IngressSamplingRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
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
		qv := convertSampling_Sflow_Interface_IngressSamplingRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate with a ONCE subscription.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/ingress-sampling-rate in the given batch object.
func (n *Sampling_Sflow_Interface_IngressSamplingRatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Interface_IngressSamplingRatePath extracts the value of the leaf IngressSamplingRate from its parent oc.Sampling_Sflow_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertSampling_Sflow_Interface_IngressSamplingRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Interface) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Interface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Interface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Interface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Interface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Interface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name with a ONCE subscription.
func (n *Sampling_Sflow_Interface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name.
func (n *Sampling_Sflow_Interface_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name in the given batch object.
func (n *Sampling_Sflow_Interface_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name.
func (n *Sampling_Sflow_Interface_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name in the given batch object.
func (n *Sampling_Sflow_Interface_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name.
func (n *Sampling_Sflow_Interface_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/name in the given batch object.
func (n *Sampling_Sflow_Interface_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Interface_NamePath extracts the value of the leaf Name from its parent oc.Sampling_Sflow_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_Interface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Interface_PollingIntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Interface_PollingIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Interface_PollingIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval with a ONCE subscription.
func (n *Sampling_Sflow_Interface_PollingIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval in the given batch object.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval in the given batch object.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/interfaces/interface/config/polling-interval in the given batch object.
func (n *Sampling_Sflow_Interface_PollingIntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Interface_PollingIntervalPath extracts the value of the leaf PollingInterval from its parent oc.Sampling_Sflow_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSampling_Sflow_Interface_PollingIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Interface) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PollingInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/polling-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_PollingIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_PollingIntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/polling-interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_PollingIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/polling-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_PollingIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_PollingIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/polling-interval with a ONCE subscription.
func (n *Sampling_Sflow_PollingIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/polling-interval.
func (n *Sampling_Sflow_PollingIntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/polling-interval in the given batch object.
func (n *Sampling_Sflow_PollingIntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/polling-interval.
func (n *Sampling_Sflow_PollingIntervalPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/polling-interval in the given batch object.
func (n *Sampling_Sflow_PollingIntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/polling-interval.
func (n *Sampling_Sflow_PollingIntervalPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/polling-interval in the given batch object.
func (n *Sampling_Sflow_PollingIntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_PollingIntervalPath extracts the value of the leaf PollingInterval from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSampling_Sflow_PollingIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PollingInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
