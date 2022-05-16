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

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, true)
	if ok {
		return convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address with a ONCE subscription.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address in the given batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address in the given batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/host-entries/host-entry/config/ipv6-address in the given batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Dns_HostEntry_Ipv6AddressPath extracts the value of the leaf Ipv6Address from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_Ipv6AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Ipv6Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/config/search with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_SearchPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns{}
	md, ok := oc.Lookup(t, n, "System_Dns", goStruct, true, true)
	if ok {
		return convertSystem_Dns_SearchPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/config/search with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_SearchPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/config/search with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_SearchPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_SearchPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/config/search with a ONCE subscription.
func (n *System_Dns_SearchPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/config/search.
func (n *System_Dns_SearchPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/config/search in the given batch object.
func (n *System_Dns_SearchPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/config/search.
func (n *System_Dns_SearchPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/config/search in the given batch object.
func (n *System_Dns_SearchPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns/config/search.
func (n *System_Dns_SearchPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/config/search in the given batch object.
func (n *System_Dns_SearchPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Dns_SearchPath extracts the value of the leaf Search from its parent oc.System_Dns
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_SearchPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Search
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_ServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns_Server {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/servers/server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_ServerPath) Get(t testing.TB) *oc.System_Dns_Server {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_ServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns_Server {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns_Server
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server with a ONCE subscription.
func (n *System_Dns_ServerPathAny) Get(t testing.TB) []*oc.System_Dns_Server {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns_Server
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/servers/server.
func (n *System_Dns_ServerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/servers/server in the given batch object.
func (n *System_Dns_ServerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/servers/server.
func (n *System_Dns_ServerPath) Replace(t testing.TB, val *oc.System_Dns_Server) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/servers/server in the given batch object.
func (n *System_Dns_ServerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Dns_Server) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns/servers/server.
func (n *System_Dns_ServerPath) Update(t testing.TB, val *oc.System_Dns_Server) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/servers/server in the given batch object.
func (n *System_Dns_ServerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Dns_Server) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server/config/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_Server_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, true, true)
	if ok {
		return convertSystem_Dns_Server_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/servers/server/config/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_Server_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server/config/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_Server_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_Server_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server/config/address with a ONCE subscription.
func (n *System_Dns_Server_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/servers/server/config/address.
func (n *System_Dns_Server_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/servers/server/config/address in the given batch object.
func (n *System_Dns_Server_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/servers/server/config/address.
func (n *System_Dns_Server_AddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/servers/server/config/address in the given batch object.
func (n *System_Dns_Server_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/dns/servers/server/config/address.
func (n *System_Dns_Server_AddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/servers/server/config/address in the given batch object.
func (n *System_Dns_Server_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Dns_Server_AddressPath extracts the value of the leaf Address from its parent oc.System_Dns_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Dns_Server_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_Server) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/dns/servers/server/config/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_Server_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, true, true)
	if ok {
		return convertSystem_Dns_Server_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-system/system/dns/servers/server/config/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_Server_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server/config/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_Server_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_Server_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server/config/port with a ONCE subscription.
func (n *System_Dns_Server_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/servers/server/config/port.
func (n *System_Dns_Server_PortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/servers/server/config/port in the given batch object.
func (n *System_Dns_Server_PortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/servers/server/config/port.
func (n *System_Dns_Server_PortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/servers/server/config/port in the given batch object.
func (n *System_Dns_Server_PortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/dns/servers/server/config/port.
func (n *System_Dns_Server_PortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/servers/server/config/port in the given batch object.
func (n *System_Dns_Server_PortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Dns_Server_PortPath extracts the value of the leaf Port from its parent oc.System_Dns_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Dns_Server_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_Server) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-system/system/config/domain-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_DomainNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, true)
	if ok {
		return convertSystem_DomainNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/config/domain-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_DomainNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/config/domain-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_DomainNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_DomainNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/config/domain-name with a ONCE subscription.
func (n *System_DomainNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/config/domain-name.
func (n *System_DomainNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/config/domain-name in the given batch object.
func (n *System_DomainNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/config/domain-name.
func (n *System_DomainNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/config/domain-name in the given batch object.
func (n *System_DomainNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/config/domain-name.
func (n *System_DomainNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/config/domain-name in the given batch object.
func (n *System_DomainNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_DomainNamePath extracts the value of the leaf DomainName from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_DomainNamePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DomainName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcServer {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServerPath) Get(t testing.TB) *oc.System_GrpcServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
func (n *System_GrpcServerPathAny) Get(t testing.TB) []*oc.System_GrpcServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GrpcServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server.
func (n *System_GrpcServerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server in the given batch object.
func (n *System_GrpcServerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server.
func (n *System_GrpcServerPath) Replace(t testing.TB, val *oc.System_GrpcServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server in the given batch object.
func (n *System_GrpcServerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_GrpcServer) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server.
func (n *System_GrpcServerPath) Update(t testing.TB, val *oc.System_GrpcServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server in the given batch object.
func (n *System_GrpcServerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_GrpcServer) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CertificateIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_CertificateIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CertificateIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CertificateIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CertificateIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id with a ONCE subscription.
func (n *System_GrpcServer_CertificateIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id.
func (n *System_GrpcServer_CertificateIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id in the given batch object.
func (n *System_GrpcServer_CertificateIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id.
func (n *System_GrpcServer_CertificateIdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id in the given batch object.
func (n *System_GrpcServer_CertificateIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id.
func (n *System_GrpcServer_CertificateIdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/certificate-id in the given batch object.
func (n *System_GrpcServer_CertificateIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_CertificateIdPath extracts the value of the leaf CertificateId from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_CertificateIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CertificateId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/enable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_EnablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_EnablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/enable with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_EnablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/enable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_EnablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_EnablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/enable with a ONCE subscription.
func (n *System_GrpcServer_EnablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/enable.
func (n *System_GrpcServer_EnablePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/enable in the given batch object.
func (n *System_GrpcServer_EnablePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/enable.
func (n *System_GrpcServer_EnablePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/enable in the given batch object.
func (n *System_GrpcServer_EnablePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/enable.
func (n *System_GrpcServer_EnablePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/enable in the given batch object.
func (n *System_GrpcServer_EnablePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_EnablePath extracts the value of the leaf Enable from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_GrpcServer_EnablePath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_ListenAddressesPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_ListenAddressesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_ListenAddressesPath) Get(t testing.TB) []oc.System_GrpcServer_ListenAddresses_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_ListenAddressesPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_ListenAddressesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses with a ONCE subscription.
func (n *System_GrpcServer_ListenAddressesPathAny) Get(t testing.TB) [][]oc.System_GrpcServer_ListenAddresses_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_GrpcServer_ListenAddresses_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses.
func (n *System_GrpcServer_ListenAddressesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses in the given batch object.
func (n *System_GrpcServer_ListenAddressesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses.
func (n *System_GrpcServer_ListenAddressesPath) Replace(t testing.TB, val []oc.System_GrpcServer_ListenAddresses_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses in the given batch object.
func (n *System_GrpcServer_ListenAddressesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.System_GrpcServer_ListenAddresses_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses.
func (n *System_GrpcServer_ListenAddressesPath) Update(t testing.TB, val []oc.System_GrpcServer_ListenAddresses_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/listen-addresses in the given batch object.
func (n *System_GrpcServer_ListenAddressesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.System_GrpcServer_ListenAddresses_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_GrpcServer_ListenAddressesPath extracts the value of the leaf ListenAddresses from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice.
func convertSystem_GrpcServer_ListenAddressesPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice{
		Metadata: md,
	}
	val := parent.ListenAddresses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_MetadataAuthenticationPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_MetadataAuthenticationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_MetadataAuthenticationPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_MetadataAuthenticationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication with a ONCE subscription.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication.
func (n *System_GrpcServer_MetadataAuthenticationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication in the given batch object.
func (n *System_GrpcServer_MetadataAuthenticationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication.
func (n *System_GrpcServer_MetadataAuthenticationPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication in the given batch object.
func (n *System_GrpcServer_MetadataAuthenticationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication.
func (n *System_GrpcServer_MetadataAuthenticationPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/metadata-authentication in the given batch object.
func (n *System_GrpcServer_MetadataAuthenticationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_MetadataAuthenticationPath extracts the value of the leaf MetadataAuthentication from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_GrpcServer_MetadataAuthenticationPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.MetadataAuthentication
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_NamePath(t, md, goStruct)
	}
	return (&oc.QualifiedString{
		Metadata: md,
	}).SetVal(goStruct.GetName())
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/name with a ONCE subscription.
func (n *System_GrpcServer_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/name.
func (n *System_GrpcServer_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/name in the given batch object.
func (n *System_GrpcServer_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/name.
func (n *System_GrpcServer_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/name in the given batch object.
func (n *System_GrpcServer_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/name.
func (n *System_GrpcServer_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/name in the given batch object.
func (n *System_GrpcServer_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_NamePath extracts the value of the leaf Name from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_NetworkInstancePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_NetworkInstancePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_NetworkInstancePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_NetworkInstancePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_NetworkInstancePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance with a ONCE subscription.
func (n *System_GrpcServer_NetworkInstancePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance.
func (n *System_GrpcServer_NetworkInstancePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance in the given batch object.
func (n *System_GrpcServer_NetworkInstancePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance.
func (n *System_GrpcServer_NetworkInstancePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance in the given batch object.
func (n *System_GrpcServer_NetworkInstancePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance.
func (n *System_GrpcServer_NetworkInstancePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/network-instance in the given batch object.
func (n *System_GrpcServer_NetworkInstancePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_NetworkInstancePath extracts the value of the leaf NetworkInstance from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_NetworkInstancePath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_PortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/port with a ONCE subscription.
func (n *System_GrpcServer_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/port.
func (n *System_GrpcServer_PortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/port in the given batch object.
func (n *System_GrpcServer_PortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/port.
func (n *System_GrpcServer_PortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/port in the given batch object.
func (n *System_GrpcServer_PortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/port.
func (n *System_GrpcServer_PortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/port in the given batch object.
func (n *System_GrpcServer_PortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_PortPath extracts the value of the leaf Port from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_GrpcServer_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/services with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_ServicesPath) Lookup(t testing.TB) *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_ServicesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/services with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_ServicesPath) Get(t testing.TB) []oc.E_SystemGrpc_GRPC_SERVICE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/services with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_ServicesPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_ServicesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/services with a ONCE subscription.
func (n *System_GrpcServer_ServicesPathAny) Get(t testing.TB) [][]oc.E_SystemGrpc_GRPC_SERVICE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_SystemGrpc_GRPC_SERVICE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/services.
func (n *System_GrpcServer_ServicesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/services in the given batch object.
func (n *System_GrpcServer_ServicesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/services.
func (n *System_GrpcServer_ServicesPath) Replace(t testing.TB, val []oc.E_SystemGrpc_GRPC_SERVICE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/services in the given batch object.
func (n *System_GrpcServer_ServicesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.E_SystemGrpc_GRPC_SERVICE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/services.
func (n *System_GrpcServer_ServicesPath) Update(t testing.TB, val []oc.E_SystemGrpc_GRPC_SERVICE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/services in the given batch object.
func (n *System_GrpcServer_ServicesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.E_SystemGrpc_GRPC_SERVICE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_GrpcServer_ServicesPath extracts the value of the leaf Services from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice.
func convertSystem_GrpcServer_ServicesPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	qv := &oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice{
		Metadata: md,
	}
	val := parent.Services
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_TransportSecurityPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, true)
	if ok {
		return convertSystem_GrpcServer_TransportSecurityPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetTransportSecurity())
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_TransportSecurityPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_TransportSecurityPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_TransportSecurityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security with a ONCE subscription.
func (n *System_GrpcServer_TransportSecurityPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security.
func (n *System_GrpcServer_TransportSecurityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security in the given batch object.
func (n *System_GrpcServer_TransportSecurityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security.
func (n *System_GrpcServer_TransportSecurityPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security in the given batch object.
func (n *System_GrpcServer_TransportSecurityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security.
func (n *System_GrpcServer_TransportSecurityPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/grpc-servers/grpc-server/config/transport-security in the given batch object.
func (n *System_GrpcServer_TransportSecurityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_GrpcServer_TransportSecurityPath extracts the value of the leaf TransportSecurity from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_GrpcServer_TransportSecurityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.TransportSecurity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/config/hostname with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_HostnamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, true)
	if ok {
		return convertSystem_HostnamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/config/hostname with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_HostnamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/config/hostname with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_HostnamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_HostnamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/config/hostname with a ONCE subscription.
func (n *System_HostnamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/config/hostname.
func (n *System_HostnamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/config/hostname in the given batch object.
func (n *System_HostnamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/config/hostname.
func (n *System_HostnamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/config/hostname in the given batch object.
func (n *System_HostnamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/config/hostname.
func (n *System_HostnamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/config/hostname in the given batch object.
func (n *System_HostnamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_HostnamePath extracts the value of the leaf Hostname from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_HostnamePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Hostname
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LicensePath) Lookup(t testing.TB) *oc.QualifiedSystem_License {
	t.Helper()
	goStruct := &oc.System_License{}
	md, ok := oc.Lookup(t, n, "System_License", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LicensePath) Get(t testing.TB) *oc.System_License {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LicensePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license with a ONCE subscription.
func (n *System_LicensePathAny) Get(t testing.TB) []*oc.System_License {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_License
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/license.
func (n *System_LicensePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/license in the given batch object.
func (n *System_LicensePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/license.
func (n *System_LicensePath) Replace(t testing.TB, val *oc.System_License) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/license in the given batch object.
func (n *System_LicensePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_License) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/license.
func (n *System_LicensePath) Update(t testing.TB, val *oc.System_License) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/license in the given batch object.
func (n *System_LicensePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_License) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_LicensePath) Lookup(t testing.TB) *oc.QualifiedSystem_License_License {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_LicensePath) Get(t testing.TB) *oc.System_License_License {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_LicensePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License_License {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License_License
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license with a ONCE subscription.
func (n *System_License_LicensePathAny) Get(t testing.TB) []*oc.System_License_License {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_License_License
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/license/licenses/license.
func (n *System_License_LicensePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/license/licenses/license in the given batch object.
func (n *System_License_LicensePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/license/licenses/license.
func (n *System_License_LicensePath) Replace(t testing.TB, val *oc.System_License_License) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/license/licenses/license in the given batch object.
func (n *System_License_LicensePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_License_License) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/license/licenses/license.
func (n *System_License_LicensePath) Update(t testing.TB, val *oc.System_License_License) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/license/licenses/license in the given batch object.
func (n *System_License_LicensePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_License_License) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/config/active with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ActivePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, true)
	if ok {
		return convertSystem_License_License_ActivePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetActive())
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/config/active with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ActivePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/config/active with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ActivePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ActivePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/config/active with a ONCE subscription.
func (n *System_License_License_ActivePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/license/licenses/license/config/active.
func (n *System_License_License_ActivePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/license/licenses/license/config/active in the given batch object.
func (n *System_License_License_ActivePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/license/licenses/license/config/active.
func (n *System_License_License_ActivePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/license/licenses/license/config/active in the given batch object.
func (n *System_License_License_ActivePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/license/licenses/license/config/active.
func (n *System_License_License_ActivePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/license/licenses/license/config/active in the given batch object.
func (n *System_License_License_ActivePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_License_License_ActivePath extracts the value of the leaf Active from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ActivePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Active
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/config/license-data with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_LicenseDataPath) Lookup(t testing.TB) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, true)
	if ok {
		return convertSystem_License_License_LicenseDataPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/config/license-data with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_LicenseDataPath) Get(t testing.TB) oc.System_License_License_LicenseData_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/config/license-data with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_LicenseDataPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License_License_LicenseData_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_LicenseDataPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/config/license-data with a ONCE subscription.
func (n *System_License_License_LicenseDataPathAny) Get(t testing.TB) []oc.System_License_License_LicenseData_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_License_License_LicenseData_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/license/licenses/license/config/license-data.
func (n *System_License_License_LicenseDataPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/license/licenses/license/config/license-data in the given batch object.
func (n *System_License_License_LicenseDataPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/license/licenses/license/config/license-data.
func (n *System_License_License_LicenseDataPath) Replace(t testing.TB, val oc.System_License_License_LicenseData_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/license/licenses/license/config/license-data in the given batch object.
func (n *System_License_License_LicenseDataPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.System_License_License_LicenseData_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/license/licenses/license/config/license-data.
func (n *System_License_License_LicenseDataPath) Update(t testing.TB, val oc.System_License_License_LicenseData_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/license/licenses/license/config/license-data in the given batch object.
func (n *System_License_License_LicenseDataPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.System_License_License_LicenseData_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_License_License_LicenseDataPath extracts the value of the leaf LicenseData from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_License_License_LicenseData_Union.
func convertSystem_License_License_LicenseDataPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_License_License_LicenseData_Union{
		Metadata: md,
	}
	val := parent.LicenseData
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/config/license-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_LicenseIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, true)
	if ok {
		return convertSystem_License_License_LicenseIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/config/license-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_LicenseIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/config/license-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_LicenseIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_LicenseIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/config/license-id with a ONCE subscription.
func (n *System_License_License_LicenseIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/license/licenses/license/config/license-id.
func (n *System_License_License_LicenseIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/license/licenses/license/config/license-id in the given batch object.
func (n *System_License_License_LicenseIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/license/licenses/license/config/license-id.
func (n *System_License_License_LicenseIdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/license/licenses/license/config/license-id in the given batch object.
func (n *System_License_License_LicenseIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/license/licenses/license/config/license-id.
func (n *System_License_License_LicenseIdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/license/licenses/license/config/license-id in the given batch object.
func (n *System_License_License_LicenseIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_License_License_LicenseIdPath extracts the value of the leaf LicenseId from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_License_License_LicenseIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LicenseId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LoggingPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging {
	t.Helper()
	goStruct := &oc.System_Logging{}
	md, ok := oc.Lookup(t, n, "System_Logging", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LoggingPath) Get(t testing.TB) *oc.System_Logging {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LoggingPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging with a ONCE subscription.
func (n *System_LoggingPathAny) Get(t testing.TB) []*oc.System_Logging {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging.
func (n *System_LoggingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging in the given batch object.
func (n *System_LoggingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging.
func (n *System_LoggingPath) Replace(t testing.TB, val *oc.System_Logging) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging in the given batch object.
func (n *System_LoggingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging.
func (n *System_LoggingPath) Update(t testing.TB, val *oc.System_Logging) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging in the given batch object.
func (n *System_LoggingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/logging/console with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_ConsolePath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_Console {
	t.Helper()
	goStruct := &oc.System_Logging_Console{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Logging_Console{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_ConsolePath) Get(t testing.TB) *oc.System_Logging_Console {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_ConsolePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_Console {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_Console
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_Console{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console with a ONCE subscription.
func (n *System_Logging_ConsolePathAny) Get(t testing.TB) []*oc.System_Logging_Console {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_Console
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/console.
func (n *System_Logging_ConsolePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/console in the given batch object.
func (n *System_Logging_ConsolePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/console.
func (n *System_Logging_ConsolePath) Replace(t testing.TB, val *oc.System_Logging_Console) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/console in the given batch object.
func (n *System_Logging_ConsolePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_Console) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/console.
func (n *System_Logging_ConsolePath) Update(t testing.TB, val *oc.System_Logging_Console) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/console in the given batch object.
func (n *System_Logging_ConsolePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_Console) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
