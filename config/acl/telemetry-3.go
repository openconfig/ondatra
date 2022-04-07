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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/traffic-class in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath extracts the value of the leaf TrafficClass from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.TrafficClass
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/config/ttl-value in the given batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_Mpls_TtlValuePath extracts the value of the leaf TtlValue from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.TtlValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_SequenceIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_SequenceIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id in the given batch object.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id in the given batch object.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/sequence-id in the given batch object.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_SequenceIdPath extracts the value of the leaf SequenceId from its parent oc.Acl_AclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_SequenceIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SequenceId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_TransportPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Transport {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_TransportPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Transport {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Transport {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Transport
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Transport {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Transport
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport.
func (n *Acl_AclSet_AclEntry_TransportPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport in the given batch object.
func (n *Acl_AclSet_AclEntry_TransportPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport.
func (n *Acl_AclSet_AclEntry_TransportPath) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_Transport) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport in the given batch object.
func (n *Acl_AclSet_AclEntry_TransportPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Transport) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport.
func (n *Acl_AclSet_AclEntry_TransportPath) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_Transport) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport in the given batch object.
func (n *Acl_AclSet_AclEntry_TransportPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Transport) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/destination-port in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Transport_DestinationPortPath extracts the value of the leaf DestinationPort from its parent oc.Acl_AclSet_AclEntry_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union.
func convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union{
		Metadata: md,
	}
	val := parent.DestinationPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Replace(t testing.TB, val oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Update(t testing.TB, val oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/source-port in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Transport_SourcePortPath extracts the value of the leaf SourcePort from its parent oc.Acl_AclSet_AclEntry_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union.
func convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union{
		Metadata: md,
	}
	val := parent.SourcePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Lookup(t testing.TB) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Get(t testing.TB) []oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Get(t testing.TB) [][]oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_PacketMatchTypes_TCP_FLAGS
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Replace(t testing.TB, val []oc.E_PacketMatchTypes_TCP_FLAGS) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.E_PacketMatchTypes_TCP_FLAGS) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Update(t testing.TB, val []oc.E_PacketMatchTypes_TCP_FLAGS) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/config/tcp-flags in the given batch object.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.E_PacketMatchTypes_TCP_FLAGS) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath extracts the value of the leaf TcpFlags from its parent oc.Acl_AclSet_AclEntry_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice.
func convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	qv := &oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice{
		Metadata: md,
	}
	val := parent.TcpFlags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/config/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/config/description with a ONCE subscription.
func (n *Acl_AclSet_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/description.
func (n *Acl_AclSet_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/config/description in the given batch object.
func (n *Acl_AclSet_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/description.
func (n *Acl_AclSet_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/config/description in the given batch object.
func (n *Acl_AclSet_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/description.
func (n *Acl_AclSet_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/config/description in the given batch object.
func (n *Acl_AclSet_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_DescriptionPath extracts the value of the leaf Description from its parent oc.Acl_AclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/config/name with a ONCE subscription.
func (n *Acl_AclSet_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/name.
func (n *Acl_AclSet_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/config/name in the given batch object.
func (n *Acl_AclSet_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/name.
func (n *Acl_AclSet_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/config/name in the given batch object.
func (n *Acl_AclSet_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/name.
func (n *Acl_AclSet_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/config/name in the given batch object.
func (n *Acl_AclSet_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_NamePath extracts the value of the leaf Name from its parent oc.Acl_AclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_TypePath) Get(t testing.TB) oc.E_Acl_ACL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/config/type with a ONCE subscription.
func (n *Acl_AclSet_TypePathAny) Get(t testing.TB) []oc.E_Acl_ACL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/type.
func (n *Acl_AclSet_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/config/type in the given batch object.
func (n *Acl_AclSet_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/type.
func (n *Acl_AclSet_TypePath) Replace(t testing.TB, val oc.E_Acl_ACL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/config/type in the given batch object.
func (n *Acl_AclSet_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_ACL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/config/type.
func (n *Acl_AclSet_TypePath) Update(t testing.TB, val oc.E_Acl_ACL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/config/type in the given batch object.
func (n *Acl_AclSet_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_ACL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_TypePath extracts the value of the leaf Type from its parent oc.Acl_AclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_TYPE.
func convertAcl_AclSet_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_InterfacePath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface {
	t.Helper()
	goStruct := &oc.Acl_Interface{}
	md, ok := oc.Lookup(t, n, "Acl_Interface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_InterfacePath) Get(t testing.TB) *oc.Acl_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
func (n *Acl_InterfacePathAny) Get(t testing.TB) []*oc.Acl_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface.
func (n *Acl_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface in the given batch object.
func (n *Acl_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface.
func (n *Acl_InterfacePath) Replace(t testing.TB, val *oc.Acl_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface in the given batch object.
func (n *Acl_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface.
func (n *Acl_InterfacePath) Update(t testing.TB, val *oc.Acl_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface in the given batch object.
func (n *Acl_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSetPath) Get(t testing.TB) *oc.Acl_Interface_EgressAclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_EgressAclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
func (n *Acl_Interface_EgressAclSetPathAny) Get(t testing.TB) []*oc.Acl_Interface_EgressAclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_EgressAclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set.
func (n *Acl_Interface_EgressAclSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set in the given batch object.
func (n *Acl_Interface_EgressAclSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set.
func (n *Acl_Interface_EgressAclSetPath) Replace(t testing.TB, val *oc.Acl_Interface_EgressAclSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set in the given batch object.
func (n *Acl_Interface_EgressAclSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface_EgressAclSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set.
func (n *Acl_Interface_EgressAclSetPath) Update(t testing.TB, val *oc.Acl_Interface_EgressAclSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set in the given batch object.
func (n *Acl_Interface_EgressAclSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface_EgressAclSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
