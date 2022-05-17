package acl

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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/destination-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv4_DscpPath extracts the value of the leaf Dscp from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Lookup(t testing.TB) *oc.QualifiedUint8Slice {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Get(t testing.TB) []uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Get(t testing.TB) [][]uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Replace(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Update(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/dscp-set in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath extracts the value of the leaf DscpSet from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8Slice.
func convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedUint8Slice {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/hop-limit in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath extracts the value of the leaf HopLimit from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/protocol in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath extracts the value of the leaf Protocol from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union.
func convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union{
		Metadata: md,
	}
	val := parent.Protocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/config/source-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Ipv6) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Ipv6) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/destination-flow-label in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath extracts the value of the leaf DestinationFlowLabel from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_DscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_DscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv6_DscpPath extracts the value of the leaf Dscp from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Ipv6_DscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) Lookup(t testing.TB) *oc.QualifiedUint8Slice {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_DscpSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) Get(t testing.TB) []uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_DscpSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPathAny) Get(t testing.TB) [][]uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) Replace(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) Update(t testing.TB, val []uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/dscp-set in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DscpSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint8) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Ipv6_DscpSetPath extracts the value of the leaf DscpSet from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8Slice.
func convertAcl_AclSet_AclEntry_Ipv6_DscpSetPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedUint8Slice {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_HopLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_HopLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/hop-limit in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_HopLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv6_HopLimitPath extracts the value of the leaf HopLimit from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Ipv6_HopLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedUint8 {
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
