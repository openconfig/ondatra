package system

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

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id.
func (n *System_Ntp_NtpKey_KeyIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id in the given batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id.
func (n *System_Ntp_NtpKey_KeyIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id in the given batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id.
func (n *System_Ntp_NtpKey_KeyIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id in the given batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_NtpKey_KeyIdPath extracts the value of the leaf KeyId from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Ntp_NtpKey_KeyIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.KeyId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyTypePath) Lookup(t testing.TB) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyTypePath) Get(t testing.TB) oc.E_System_NTP_AUTH_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_System_NTP_AUTH_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Get(t testing.TB) []oc.E_System_NTP_AUTH_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_System_NTP_AUTH_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type.
func (n *System_Ntp_NtpKey_KeyTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type in the given batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type.
func (n *System_Ntp_NtpKey_KeyTypePath) Replace(t testing.TB, val oc.E_System_NTP_AUTH_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type in the given batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_System_NTP_AUTH_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type.
func (n *System_Ntp_NtpKey_KeyTypePath) Update(t testing.TB, val oc.E_System_NTP_AUTH_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type in the given batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_System_NTP_AUTH_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Ntp_NtpKey_KeyTypePath extracts the value of the leaf KeyType from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedE_System_NTP_AUTH_TYPE.
func convertSystem_Ntp_NtpKey_KeyTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_System_NTP_AUTH_TYPE{
		Metadata: md,
	}
	val := parent.KeyType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyValuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyValuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value.
func (n *System_Ntp_NtpKey_KeyValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value in the given batch object.
func (n *System_Ntp_NtpKey_KeyValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value.
func (n *System_Ntp_NtpKey_KeyValuePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value in the given batch object.
func (n *System_Ntp_NtpKey_KeyValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value.
func (n *System_Ntp_NtpKey_KeyValuePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-value in the given batch object.
func (n *System_Ntp_NtpKey_KeyValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_NtpKey_KeyValuePath extracts the value of the leaf KeyValue from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Ntp_NtpKey_KeyValuePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.KeyValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/config/ntp-source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpSourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_NtpSourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/config/ntp-source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpSourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/config/ntp-source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpSourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpSourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/config/ntp-source-address with a ONCE subscription.
func (n *System_Ntp_NtpSourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/config/ntp-source-address.
func (n *System_Ntp_NtpSourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/config/ntp-source-address in the given batch object.
func (n *System_Ntp_NtpSourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/config/ntp-source-address.
func (n *System_Ntp_NtpSourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/config/ntp-source-address in the given batch object.
func (n *System_Ntp_NtpSourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/config/ntp-source-address.
func (n *System_Ntp_NtpSourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/config/ntp-source-address in the given batch object.
func (n *System_Ntp_NtpSourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_NtpSourceAddressPath extracts the value of the leaf NtpSourceAddress from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Ntp_NtpSourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NtpSourceAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_ServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Ntp_Server {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Ntp_Server{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_ServerPath) Get(t testing.TB) *oc.System_Ntp_Server {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_ServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Ntp_Server {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Ntp_Server
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Ntp_Server{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server with a ONCE subscription.
func (n *System_Ntp_ServerPathAny) Get(t testing.TB) []*oc.System_Ntp_Server {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Ntp_Server
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server.
func (n *System_Ntp_ServerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server in the given batch object.
func (n *System_Ntp_ServerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server.
func (n *System_Ntp_ServerPath) Replace(t testing.TB, val *oc.System_Ntp_Server) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server in the given batch object.
func (n *System_Ntp_ServerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Ntp_Server) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server.
func (n *System_Ntp_ServerPath) Update(t testing.TB, val *oc.System_Ntp_Server) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server in the given batch object.
func (n *System_Ntp_ServerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Ntp_Server) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server/config/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_Server_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_Server_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server/config/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_Server_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server/config/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_Server_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_Server_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server/config/address with a ONCE subscription.
func (n *System_Ntp_Server_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server/config/address.
func (n *System_Ntp_Server_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server/config/address in the given batch object.
func (n *System_Ntp_Server_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server/config/address.
func (n *System_Ntp_Server_AddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server/config/address in the given batch object.
func (n *System_Ntp_Server_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server/config/address.
func (n *System_Ntp_Server_AddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server/config/address in the given batch object.
func (n *System_Ntp_Server_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_Server_AddressPath extracts the value of the leaf Address from its parent oc.System_Ntp_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Ntp_Server_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_Server) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server/config/association-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_Server_AssociationTypePath) Lookup(t testing.TB) *oc.QualifiedE_Server_AssociationType {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_Server_AssociationTypePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Server_AssociationType{
		Metadata: md,
	}).SetVal(goStruct.GetAssociationType())
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server/config/association-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_Server_AssociationTypePath) Get(t testing.TB) oc.E_Server_AssociationType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server/config/association-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_Server_AssociationTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Server_AssociationType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Server_AssociationType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_Server_AssociationTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server/config/association-type with a ONCE subscription.
func (n *System_Ntp_Server_AssociationTypePathAny) Get(t testing.TB) []oc.E_Server_AssociationType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Server_AssociationType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server/config/association-type.
func (n *System_Ntp_Server_AssociationTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server/config/association-type in the given batch object.
func (n *System_Ntp_Server_AssociationTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server/config/association-type.
func (n *System_Ntp_Server_AssociationTypePath) Replace(t testing.TB, val oc.E_Server_AssociationType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server/config/association-type in the given batch object.
func (n *System_Ntp_Server_AssociationTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Server_AssociationType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server/config/association-type.
func (n *System_Ntp_Server_AssociationTypePath) Update(t testing.TB, val oc.E_Server_AssociationType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server/config/association-type in the given batch object.
func (n *System_Ntp_Server_AssociationTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Server_AssociationType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Ntp_Server_AssociationTypePath extracts the value of the leaf AssociationType from its parent oc.System_Ntp_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Server_AssociationType.
func convertSystem_Ntp_Server_AssociationTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_Server) *oc.QualifiedE_Server_AssociationType {
	t.Helper()
	qv := &oc.QualifiedE_Server_AssociationType{
		Metadata: md,
	}
	val := parent.AssociationType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server/config/iburst with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_Server_IburstPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_Server_IburstPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetIburst())
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server/config/iburst with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_Server_IburstPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server/config/iburst with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_Server_IburstPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_Server_IburstPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server/config/iburst with a ONCE subscription.
func (n *System_Ntp_Server_IburstPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server/config/iburst.
func (n *System_Ntp_Server_IburstPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server/config/iburst in the given batch object.
func (n *System_Ntp_Server_IburstPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server/config/iburst.
func (n *System_Ntp_Server_IburstPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server/config/iburst in the given batch object.
func (n *System_Ntp_Server_IburstPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server/config/iburst.
func (n *System_Ntp_Server_IburstPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server/config/iburst in the given batch object.
func (n *System_Ntp_Server_IburstPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_Server_IburstPath extracts the value of the leaf Iburst from its parent oc.System_Ntp_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Ntp_Server_IburstPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_Server) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Iburst
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server/config/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_Server_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_Server_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server/config/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_Server_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server/config/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_Server_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_Server_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server/config/port with a ONCE subscription.
func (n *System_Ntp_Server_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server/config/port.
func (n *System_Ntp_Server_PortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server/config/port in the given batch object.
func (n *System_Ntp_Server_PortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server/config/port.
func (n *System_Ntp_Server_PortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server/config/port in the given batch object.
func (n *System_Ntp_Server_PortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server/config/port.
func (n *System_Ntp_Server_PortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server/config/port in the given batch object.
func (n *System_Ntp_Server_PortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_Server_PortPath extracts the value of the leaf Port from its parent oc.System_Ntp_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Ntp_Server_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_Server) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server/config/prefer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_Server_PreferPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_Server_PreferPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPrefer())
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server/config/prefer with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_Server_PreferPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server/config/prefer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_Server_PreferPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_Server_PreferPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server/config/prefer with a ONCE subscription.
func (n *System_Ntp_Server_PreferPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server/config/prefer.
func (n *System_Ntp_Server_PreferPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server/config/prefer in the given batch object.
func (n *System_Ntp_Server_PreferPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server/config/prefer.
func (n *System_Ntp_Server_PreferPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server/config/prefer in the given batch object.
func (n *System_Ntp_Server_PreferPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server/config/prefer.
func (n *System_Ntp_Server_PreferPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server/config/prefer in the given batch object.
func (n *System_Ntp_Server_PreferPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_Server_PreferPath extracts the value of the leaf Prefer from its parent oc.System_Ntp_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Ntp_Server_PreferPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_Server) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Prefer
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/servers/server/config/version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_Server_VersionPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Ntp_Server{}
	md, ok := oc.Lookup(t, n, "System_Ntp_Server", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_Server_VersionPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetVersion())
}

// Get fetches the value at /openconfig-system/system/ntp/servers/server/config/version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_Server_VersionPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/servers/server/config/version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_Server_VersionPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_Server_VersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/servers/server/config/version with a ONCE subscription.
func (n *System_Ntp_Server_VersionPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/servers/server/config/version.
func (n *System_Ntp_Server_VersionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/servers/server/config/version in the given batch object.
func (n *System_Ntp_Server_VersionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/servers/server/config/version.
func (n *System_Ntp_Server_VersionPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/servers/server/config/version in the given batch object.
func (n *System_Ntp_Server_VersionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/servers/server/config/version.
func (n *System_Ntp_Server_VersionPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/servers/server/config/version in the given batch object.
func (n *System_Ntp_Server_VersionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_Server_VersionPath extracts the value of the leaf Version from its parent oc.System_Ntp_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Ntp_Server_VersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_Server) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Version
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_SshServer {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_SshServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServerPath) Get(t testing.TB) *oc.System_SshServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_SshServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_SshServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_SshServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server with a ONCE subscription.
func (n *System_SshServerPathAny) Get(t testing.TB) []*oc.System_SshServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_SshServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ssh-server.
func (n *System_SshServerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ssh-server in the given batch object.
func (n *System_SshServerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ssh-server.
func (n *System_SshServerPath) Replace(t testing.TB, val *oc.System_SshServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ssh-server in the given batch object.
func (n *System_SshServerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_SshServer) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ssh-server.
func (n *System_SshServerPath) Update(t testing.TB, val *oc.System_SshServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ssh-server in the given batch object.
func (n *System_SshServerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_SshServer) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/config/enable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_EnablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, true)
	if ok {
		return convertSystem_SshServer_EnablePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnable())
}

// Get fetches the value at /openconfig-system/system/ssh-server/config/enable with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_EnablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/config/enable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_EnablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_EnablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/config/enable with a ONCE subscription.
func (n *System_SshServer_EnablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ssh-server/config/enable.
func (n *System_SshServer_EnablePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ssh-server/config/enable in the given batch object.
func (n *System_SshServer_EnablePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ssh-server/config/enable.
func (n *System_SshServer_EnablePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ssh-server/config/enable in the given batch object.
func (n *System_SshServer_EnablePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ssh-server/config/enable.
func (n *System_SshServer_EnablePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ssh-server/config/enable in the given batch object.
func (n *System_SshServer_EnablePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_SshServer_EnablePath extracts the value of the leaf Enable from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_SshServer_EnablePath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/config/protocol-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_ProtocolVersionPath) Lookup(t testing.TB) *oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, true)
	if ok {
		return convertSystem_SshServer_ProtocolVersionPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_SshServer_ProtocolVersion{
		Metadata: md,
	}).SetVal(goStruct.GetProtocolVersion())
}

// Get fetches the value at /openconfig-system/system/ssh-server/config/protocol-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_ProtocolVersionPath) Get(t testing.TB) oc.E_SshServer_ProtocolVersion {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/config/protocol-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_ProtocolVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SshServer_ProtocolVersion
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_ProtocolVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/config/protocol-version with a ONCE subscription.
func (n *System_SshServer_ProtocolVersionPathAny) Get(t testing.TB) []oc.E_SshServer_ProtocolVersion {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SshServer_ProtocolVersion
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ssh-server/config/protocol-version.
func (n *System_SshServer_ProtocolVersionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ssh-server/config/protocol-version in the given batch object.
func (n *System_SshServer_ProtocolVersionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ssh-server/config/protocol-version.
func (n *System_SshServer_ProtocolVersionPath) Replace(t testing.TB, val oc.E_SshServer_ProtocolVersion) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ssh-server/config/protocol-version in the given batch object.
func (n *System_SshServer_ProtocolVersionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SshServer_ProtocolVersion) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ssh-server/config/protocol-version.
func (n *System_SshServer_ProtocolVersionPath) Update(t testing.TB, val oc.E_SshServer_ProtocolVersion) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ssh-server/config/protocol-version in the given batch object.
func (n *System_SshServer_ProtocolVersionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SshServer_ProtocolVersion) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_SshServer_ProtocolVersionPath extracts the value of the leaf ProtocolVersion from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SshServer_ProtocolVersion.
func convertSystem_SshServer_ProtocolVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	qv := &oc.QualifiedE_SshServer_ProtocolVersion{
		Metadata: md,
	}
	val := parent.ProtocolVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/config/rate-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_RateLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, true)
	if ok {
		return convertSystem_SshServer_RateLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server/config/rate-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_RateLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/config/rate-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_RateLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_RateLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/config/rate-limit with a ONCE subscription.
func (n *System_SshServer_RateLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ssh-server/config/rate-limit.
func (n *System_SshServer_RateLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ssh-server/config/rate-limit in the given batch object.
func (n *System_SshServer_RateLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ssh-server/config/rate-limit.
func (n *System_SshServer_RateLimitPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ssh-server/config/rate-limit in the given batch object.
func (n *System_SshServer_RateLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ssh-server/config/rate-limit.
func (n *System_SshServer_RateLimitPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ssh-server/config/rate-limit in the given batch object.
func (n *System_SshServer_RateLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_SshServer_RateLimitPath extracts the value of the leaf RateLimit from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_SshServer_RateLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RateLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/config/session-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_SessionLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, true)
	if ok {
		return convertSystem_SshServer_SessionLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server/config/session-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_SessionLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/config/session-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_SessionLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_SessionLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/config/session-limit with a ONCE subscription.
func (n *System_SshServer_SessionLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ssh-server/config/session-limit.
func (n *System_SshServer_SessionLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ssh-server/config/session-limit in the given batch object.
func (n *System_SshServer_SessionLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ssh-server/config/session-limit.
func (n *System_SshServer_SessionLimitPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ssh-server/config/session-limit in the given batch object.
func (n *System_SshServer_SessionLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ssh-server/config/session-limit.
func (n *System_SshServer_SessionLimitPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ssh-server/config/session-limit in the given batch object.
func (n *System_SshServer_SessionLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_SshServer_SessionLimitPath extracts the value of the leaf SessionLimit from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_SshServer_SessionLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SessionLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/config/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, true)
	if ok {
		return convertSystem_SshServer_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server/config/timeout with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_TimeoutPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/config/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/config/timeout with a ONCE subscription.
func (n *System_SshServer_TimeoutPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ssh-server/config/timeout.
func (n *System_SshServer_TimeoutPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ssh-server/config/timeout in the given batch object.
func (n *System_SshServer_TimeoutPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ssh-server/config/timeout.
func (n *System_SshServer_TimeoutPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ssh-server/config/timeout in the given batch object.
func (n *System_SshServer_TimeoutPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ssh-server/config/timeout.
func (n *System_SshServer_TimeoutPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ssh-server/config/timeout in the given batch object.
func (n *System_SshServer_TimeoutPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_SshServer_TimeoutPath extracts the value of the leaf Timeout from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_SshServer_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_TelnetServer {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_TelnetServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServerPath) Get(t testing.TB) *oc.System_TelnetServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_TelnetServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_TelnetServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_TelnetServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server with a ONCE subscription.
func (n *System_TelnetServerPathAny) Get(t testing.TB) []*oc.System_TelnetServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_TelnetServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/telnet-server.
func (n *System_TelnetServerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/telnet-server in the given batch object.
func (n *System_TelnetServerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/telnet-server.
func (n *System_TelnetServerPath) Replace(t testing.TB, val *oc.System_TelnetServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/telnet-server in the given batch object.
func (n *System_TelnetServerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_TelnetServer) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/telnet-server.
func (n *System_TelnetServerPath) Update(t testing.TB, val *oc.System_TelnetServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/telnet-server in the given batch object.
func (n *System_TelnetServerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_TelnetServer) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/config/enable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_EnablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, true)
	if ok {
		return convertSystem_TelnetServer_EnablePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnable())
}

// Get fetches the value at /openconfig-system/system/telnet-server/config/enable with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_EnablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/config/enable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_EnablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_EnablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/config/enable with a ONCE subscription.
func (n *System_TelnetServer_EnablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/telnet-server/config/enable.
func (n *System_TelnetServer_EnablePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/telnet-server/config/enable in the given batch object.
func (n *System_TelnetServer_EnablePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/telnet-server/config/enable.
func (n *System_TelnetServer_EnablePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/telnet-server/config/enable in the given batch object.
func (n *System_TelnetServer_EnablePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/telnet-server/config/enable.
func (n *System_TelnetServer_EnablePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/telnet-server/config/enable in the given batch object.
func (n *System_TelnetServer_EnablePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_TelnetServer_EnablePath extracts the value of the leaf Enable from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_TelnetServer_EnablePath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/config/rate-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_RateLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, true)
	if ok {
		return convertSystem_TelnetServer_RateLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server/config/rate-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_RateLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/config/rate-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_RateLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_RateLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/config/rate-limit with a ONCE subscription.
func (n *System_TelnetServer_RateLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/telnet-server/config/rate-limit.
func (n *System_TelnetServer_RateLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/telnet-server/config/rate-limit in the given batch object.
func (n *System_TelnetServer_RateLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/telnet-server/config/rate-limit.
func (n *System_TelnetServer_RateLimitPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/telnet-server/config/rate-limit in the given batch object.
func (n *System_TelnetServer_RateLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/telnet-server/config/rate-limit.
func (n *System_TelnetServer_RateLimitPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/telnet-server/config/rate-limit in the given batch object.
func (n *System_TelnetServer_RateLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_TelnetServer_RateLimitPath extracts the value of the leaf RateLimit from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_TelnetServer_RateLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RateLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/config/session-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_SessionLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, true)
	if ok {
		return convertSystem_TelnetServer_SessionLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server/config/session-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_SessionLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/config/session-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_SessionLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_SessionLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/config/session-limit with a ONCE subscription.
func (n *System_TelnetServer_SessionLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/telnet-server/config/session-limit.
func (n *System_TelnetServer_SessionLimitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/telnet-server/config/session-limit in the given batch object.
func (n *System_TelnetServer_SessionLimitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/telnet-server/config/session-limit.
func (n *System_TelnetServer_SessionLimitPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/telnet-server/config/session-limit in the given batch object.
func (n *System_TelnetServer_SessionLimitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/telnet-server/config/session-limit.
func (n *System_TelnetServer_SessionLimitPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/telnet-server/config/session-limit in the given batch object.
func (n *System_TelnetServer_SessionLimitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_TelnetServer_SessionLimitPath extracts the value of the leaf SessionLimit from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_TelnetServer_SessionLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SessionLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/config/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, true)
	if ok {
		return convertSystem_TelnetServer_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server/config/timeout with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_TimeoutPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/config/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/config/timeout with a ONCE subscription.
func (n *System_TelnetServer_TimeoutPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/telnet-server/config/timeout.
func (n *System_TelnetServer_TimeoutPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/telnet-server/config/timeout in the given batch object.
func (n *System_TelnetServer_TimeoutPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/telnet-server/config/timeout.
func (n *System_TelnetServer_TimeoutPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/telnet-server/config/timeout in the given batch object.
func (n *System_TelnetServer_TimeoutPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/telnet-server/config/timeout.
func (n *System_TelnetServer_TimeoutPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/telnet-server/config/timeout in the given batch object.
func (n *System_TelnetServer_TimeoutPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_TelnetServer_TimeoutPath extracts the value of the leaf Timeout from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_TelnetServer_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
