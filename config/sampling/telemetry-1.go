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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Collector_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/address with a ONCE subscription.
func (n *Sampling_Sflow_Collector_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/address.
func (n *Sampling_Sflow_Collector_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/address in the given batch object.
func (n *Sampling_Sflow_Collector_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/address.
func (n *Sampling_Sflow_Collector_AddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/address in the given batch object.
func (n *Sampling_Sflow_Collector_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/address.
func (n *Sampling_Sflow_Collector_AddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/address in the given batch object.
func (n *Sampling_Sflow_Collector_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Collector_AddressPath extracts the value of the leaf Address from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_Collector_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Collector_NetworkInstancePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_NetworkInstancePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance with a ONCE subscription.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance in the given batch object.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance in the given batch object.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/network-instance in the given batch object.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Collector_NetworkInstancePath extracts the value of the leaf NetworkInstance from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_Collector_NetworkInstancePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NetworkInstance
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Collector_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/port with a ONCE subscription.
func (n *Sampling_Sflow_Collector_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/port.
func (n *Sampling_Sflow_Collector_PortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/port in the given batch object.
func (n *Sampling_Sflow_Collector_PortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/port.
func (n *Sampling_Sflow_Collector_PortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/port in the given batch object.
func (n *Sampling_Sflow_Collector_PortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/port.
func (n *Sampling_Sflow_Collector_PortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/port in the given batch object.
func (n *Sampling_Sflow_Collector_PortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Collector_PortPath extracts the value of the leaf Port from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSampling_Sflow_Collector_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Port
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_Collector_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address with a ONCE subscription.
func (n *Sampling_Sflow_Collector_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address.
func (n *Sampling_Sflow_Collector_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address in the given batch object.
func (n *Sampling_Sflow_Collector_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address.
func (n *Sampling_Sflow_Collector_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address in the given batch object.
func (n *Sampling_Sflow_Collector_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address.
func (n *Sampling_Sflow_Collector_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/collectors/collector/config/source-address in the given batch object.
func (n *Sampling_Sflow_Collector_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_Collector_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_Collector_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/config/dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_DscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, true)
	if ok {
		return convertSampling_Sflow_DscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/config/dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_DscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/config/dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_DscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_DscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/config/dscp with a ONCE subscription.
func (n *Sampling_Sflow_DscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-sampling/sampling/sflow/config/dscp.
func (n *Sampling_Sflow_DscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-sampling/sampling/sflow/config/dscp in the given batch object.
func (n *Sampling_Sflow_DscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-sampling/sampling/sflow/config/dscp.
func (n *Sampling_Sflow_DscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-sampling/sampling/sflow/config/dscp in the given batch object.
func (n *Sampling_Sflow_DscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-sampling/sampling/sflow/config/dscp.
func (n *Sampling_Sflow_DscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-sampling/sampling/sflow/config/dscp in the given batch object.
func (n *Sampling_Sflow_DscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSampling_Sflow_DscpPath extracts the value of the leaf Dscp from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSampling_Sflow_DscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedUint8 {
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
