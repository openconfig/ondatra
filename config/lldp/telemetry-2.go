package lldp

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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Lldp_Interface{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface", goStruct, true, true)
	if ok {
		return convertLldp_Interface_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/config/enabled with a ONCE subscription.
func (n *Lldp_Interface_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lldp/lldp/interfaces/interface/config/enabled.
func (n *Lldp_Interface_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lldp/lldp/interfaces/interface/config/enabled in the given batch object.
func (n *Lldp_Interface_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lldp/lldp/interfaces/interface/config/enabled.
func (n *Lldp_Interface_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-lldp/lldp/interfaces/interface/config/enabled in the given batch object.
func (n *Lldp_Interface_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-lldp/lldp/interfaces/interface/config/enabled.
func (n *Lldp_Interface_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-lldp/lldp/interfaces/interface/config/enabled in the given batch object.
func (n *Lldp_Interface_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLldp_Interface_EnabledPath extracts the value of the leaf Enabled from its parent oc.Lldp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLldp_Interface_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface", goStruct, true, true)
	if ok {
		return convertLldp_Interface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/config/name with a ONCE subscription.
func (n *Lldp_Interface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lldp/lldp/interfaces/interface/config/name.
func (n *Lldp_Interface_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lldp/lldp/interfaces/interface/config/name in the given batch object.
func (n *Lldp_Interface_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lldp/lldp/interfaces/interface/config/name.
func (n *Lldp_Interface_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-lldp/lldp/interfaces/interface/config/name in the given batch object.
func (n *Lldp_Interface_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-lldp/lldp/interfaces/interface/config/name.
func (n *Lldp_Interface_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-lldp/lldp/interfaces/interface/config/name in the given batch object.
func (n *Lldp_Interface_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLldp_Interface_NamePath extracts the value of the leaf Name from its parent oc.Lldp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-lldp/lldp/config/suppress-tlv-advertisement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_SuppressTlvAdvertisementPath) Lookup(t testing.TB) *oc.QualifiedE_LldpTypes_LLDP_TLVSlice {
	t.Helper()
	goStruct := &oc.Lldp{}
	md, ok := oc.Lookup(t, n, "Lldp", goStruct, true, true)
	if ok {
		return convertLldp_SuppressTlvAdvertisementPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/config/suppress-tlv-advertisement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_SuppressTlvAdvertisementPath) Get(t testing.TB) []oc.E_LldpTypes_LLDP_TLV {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/config/suppress-tlv-advertisement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_SuppressTlvAdvertisementPathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpTypes_LLDP_TLVSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpTypes_LLDP_TLVSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLldp_SuppressTlvAdvertisementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/config/suppress-tlv-advertisement with a ONCE subscription.
func (n *Lldp_SuppressTlvAdvertisementPathAny) Get(t testing.TB) [][]oc.E_LldpTypes_LLDP_TLV {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_LldpTypes_LLDP_TLV
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lldp/lldp/config/suppress-tlv-advertisement.
func (n *Lldp_SuppressTlvAdvertisementPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lldp/lldp/config/suppress-tlv-advertisement in the given batch object.
func (n *Lldp_SuppressTlvAdvertisementPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lldp/lldp/config/suppress-tlv-advertisement.
func (n *Lldp_SuppressTlvAdvertisementPath) Replace(t testing.TB, val []oc.E_LldpTypes_LLDP_TLV) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-lldp/lldp/config/suppress-tlv-advertisement in the given batch object.
func (n *Lldp_SuppressTlvAdvertisementPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.E_LldpTypes_LLDP_TLV) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-lldp/lldp/config/suppress-tlv-advertisement.
func (n *Lldp_SuppressTlvAdvertisementPath) Update(t testing.TB, val []oc.E_LldpTypes_LLDP_TLV) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-lldp/lldp/config/suppress-tlv-advertisement in the given batch object.
func (n *Lldp_SuppressTlvAdvertisementPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.E_LldpTypes_LLDP_TLV) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertLldp_SuppressTlvAdvertisementPath extracts the value of the leaf SuppressTlvAdvertisement from its parent oc.Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpTypes_LLDP_TLVSlice.
func convertLldp_SuppressTlvAdvertisementPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp) *oc.QualifiedE_LldpTypes_LLDP_TLVSlice {
	t.Helper()
	qv := &oc.QualifiedE_LldpTypes_LLDP_TLVSlice{
		Metadata: md,
	}
	val := parent.SuppressTlvAdvertisement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
