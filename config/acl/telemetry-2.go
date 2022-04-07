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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_ProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_ProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/protocol in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_ProtocolPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Ipv6_Protocol_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Ipv6_ProtocolPath extracts the value of the leaf Protocol from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union.
func convertAcl_AclSet_AclEntry_Ipv6_ProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union{
		Metadata: md,
	}
	val := parent.Protocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertAcl_AclSet_AclEntry_Ipv6_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-address in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv6_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv6_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
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
		qv := convertAcl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/config/source-flow-label in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath extracts the value of the leaf SourceFlowLabel from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_Ipv6_SourceFlowLabelPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SourceFlowLabel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2Path) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_L2 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_L2{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2Path) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_L2 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2PathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_L2 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_L2
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_L2{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2PathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_L2 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_L2
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2.
func (n *Acl_AclSet_AclEntry_L2Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 in the given batch object.
func (n *Acl_AclSet_AclEntry_L2Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2.
func (n *Acl_AclSet_AclEntry_L2Path) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_L2) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 in the given batch object.
func (n *Acl_AclSet_AclEntry_L2Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_L2) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2.
func (n *Acl_AclSet_AclEntry_L2Path) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_L2) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2 in the given batch object.
func (n *Acl_AclSet_AclEntry_L2Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_L2) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac-mask in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath extracts the value of the leaf DestinationMacMask from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/destination-mac in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_L2_DestinationMacPath extracts the value of the leaf DestinationMac from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_EthertypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_EthertypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_L2_Ethertype_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_L2_Ethertype_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_L2_Ethertype_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_L2_Ethertype_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/ethertype in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_L2_Ethertype_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_L2_EthertypePath extracts the value of the leaf Ethertype from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union.
func convertAcl_AclSet_AclEntry_L2_EthertypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union{
		Metadata: md,
	}
	val := parent.Ethertype
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac-mask in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath extracts the value of the leaf SourceMacMask from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_SourceMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_SourceMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/config/source-mac in the given batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_L2_SourceMacPath extracts the value of the leaf SourceMac from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_SourceMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_MplsPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_MplsPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Mpls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Mpls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Mpls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Mpls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls.
func (n *Acl_AclSet_AclEntry_MplsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls in the given batch object.
func (n *Acl_AclSet_AclEntry_MplsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls.
func (n *Acl_AclSet_AclEntry_MplsPath) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_Mpls) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls in the given batch object.
func (n *Acl_AclSet_AclEntry_MplsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Mpls) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls.
func (n *Acl_AclSet_AclEntry_MplsPath) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_Mpls) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls in the given batch object.
func (n *Acl_AclSet_AclEntry_MplsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Mpls) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/end-label-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath extracts the value of the leaf EndLabelValue from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union.
func convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union{
		Metadata: md,
	}
	val := parent.EndLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/start-label-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath extracts the value of the leaf StartLabelValue from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union.
func convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union{
		Metadata: md,
	}
	val := parent.StartLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
