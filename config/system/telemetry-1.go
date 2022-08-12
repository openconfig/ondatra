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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup_Server_Radius
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius in the given batch object.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Replace(t testing.TB, val *oc.System_Aaa_ServerGroup_Server_Radius) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius in the given batch object.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_ServerGroup_Server_Radius) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Update(t testing.TB, val *oc.System_Aaa_ServerGroup_Server_Radius) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius in the given batch object.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_ServerGroup_Server_Radius) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAcctPort())
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/acct-port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath extracts the value of the leaf AcctPort from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AcctPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAuthPort())
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/auth-port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath extracts the value of the leaf AuthPort from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AuthPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/retransmit-attempts in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath extracts the value of the leaf RetransmitAttempts from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.RetransmitAttempts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key-hashed in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath extracts the value of the leaf SecretKeyHashed from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SecretKeyHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/secret-key in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath extracts the value of the leaf SecretKey from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SecretKey
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/config/source-address in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup_Server_Tacacs
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs in the given batch object.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Replace(t testing.TB, val *oc.System_Aaa_ServerGroup_Server_Tacacs) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs in the given batch object.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_ServerGroup_Server_Tacacs) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Update(t testing.TB, val *oc.System_Aaa_ServerGroup_Server_Tacacs) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs in the given batch object.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_ServerGroup_Server_Tacacs) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/port in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath extracts the value of the leaf Port from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key-hashed in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath extracts the value of the leaf SecretKeyHashed from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SecretKeyHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/secret-key in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath extracts the value of the leaf SecretKey from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SecretKey
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/config/source-address in the given batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout in the given batch object.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout in the given batch object.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/servers/server/config/timeout in the given batch object.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_ServerGroup_Server_TimeoutPath extracts the value of the leaf Timeout from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_TypePath) Lookup(t testing.TB) *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_ServerGroup_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_TypePath) Get(t testing.TB) oc.E_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/config/type with a ONCE subscription.
func (n *System_Aaa_ServerGroup_TypePathAny) Get(t testing.TB) []oc.E_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AaaTypes_AAA_SERVER_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/server-groups/server-group/config/type.
func (n *System_Aaa_ServerGroup_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/server-groups/server-group/config/type in the given batch object.
func (n *System_Aaa_ServerGroup_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/server-groups/server-group/config/type.
func (n *System_Aaa_ServerGroup_TypePath) Replace(t testing.TB, val oc.E_AaaTypes_AAA_SERVER_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/server-groups/server-group/config/type in the given batch object.
func (n *System_Aaa_ServerGroup_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_AaaTypes_AAA_SERVER_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/server-groups/server-group/config/type.
func (n *System_Aaa_ServerGroup_TypePath) Update(t testing.TB, val oc.E_AaaTypes_AAA_SERVER_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/server-groups/server-group/config/type in the given batch object.
func (n *System_Aaa_ServerGroup_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_AaaTypes_AAA_SERVER_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_ServerGroup_TypePath extracts the value of the leaf Type from its parent oc.System_Aaa_ServerGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE.
func convertSystem_Aaa_ServerGroup_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup) *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/clock with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_ClockPath) Lookup(t testing.TB) *oc.QualifiedSystem_Clock {
	t.Helper()
	goStruct := &oc.System_Clock{}
	md, ok := oc.Lookup(t, n, "System_Clock", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Clock{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/clock with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_ClockPath) Get(t testing.TB) *oc.System_Clock {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/clock with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_ClockPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Clock {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Clock
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Clock{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Clock", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Clock{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/clock with a ONCE subscription.
func (n *System_ClockPathAny) Get(t testing.TB) []*oc.System_Clock {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Clock
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/clock.
func (n *System_ClockPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/clock in the given batch object.
func (n *System_ClockPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/clock.
func (n *System_ClockPath) Replace(t testing.TB, val *oc.System_Clock) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/clock in the given batch object.
func (n *System_ClockPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Clock) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/clock.
func (n *System_ClockPath) Update(t testing.TB, val *oc.System_Clock) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/clock in the given batch object.
func (n *System_ClockPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Clock) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/clock/config/timezone-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Clock_TimezoneNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Clock{}
	md, ok := oc.Lookup(t, n, "System_Clock", goStruct, true, true)
	if ok {
		return convertSystem_Clock_TimezoneNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/clock/config/timezone-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Clock_TimezoneNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/clock/config/timezone-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Clock_TimezoneNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Clock{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Clock", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Clock_TimezoneNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/clock/config/timezone-name with a ONCE subscription.
func (n *System_Clock_TimezoneNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/clock/config/timezone-name.
func (n *System_Clock_TimezoneNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/clock/config/timezone-name in the given batch object.
func (n *System_Clock_TimezoneNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/clock/config/timezone-name.
func (n *System_Clock_TimezoneNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/clock/config/timezone-name in the given batch object.
func (n *System_Clock_TimezoneNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/clock/config/timezone-name.
func (n *System_Clock_TimezoneNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/clock/config/timezone-name in the given batch object.
func (n *System_Clock_TimezoneNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Clock_TimezoneNamePath extracts the value of the leaf TimezoneName from its parent oc.System_Clock
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Clock_TimezoneNamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Clock) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TimezoneName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_DnsPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns {
	t.Helper()
	goStruct := &oc.System_Dns{}
	md, ok := oc.Lookup(t, n, "System_Dns", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Dns{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_DnsPath) Get(t testing.TB) *oc.System_Dns {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_DnsPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns with a ONCE subscription.
func (n *System_DnsPathAny) Get(t testing.TB) []*oc.System_Dns {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns.
func (n *System_DnsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns in the given batch object.
func (n *System_DnsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns.
func (n *System_DnsPath) Replace(t testing.TB, val *oc.System_Dns) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns in the given batch object.
func (n *System_DnsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Dns) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns.
func (n *System_DnsPath) Update(t testing.TB, val *oc.System_Dns) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns in the given batch object.
func (n *System_DnsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Dns) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntryPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns_HostEntry {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Dns_HostEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntryPath) Get(t testing.TB) *oc.System_Dns_HostEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns_HostEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns_HostEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns_HostEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription.
func (n *System_Dns_HostEntryPathAny) Get(t testing.TB) []*oc.System_Dns_HostEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns_HostEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/host-entries/host-entry.
func (n *System_Dns_HostEntryPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/host-entries/host-entry in the given batch object.
func (n *System_Dns_HostEntryPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/host-entries/host-entry.
func (n *System_Dns_HostEntryPath) Replace(t testing.TB, val *oc.System_Dns_HostEntry) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/host-entries/host-entry in the given batch object.
func (n *System_Dns_HostEntryPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Dns_HostEntry) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns/host-entries/host-entry.
func (n *System_Dns_HostEntryPath) Update(t testing.TB, val *oc.System_Dns_HostEntry) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/host-entries/host-entry in the given batch object.
func (n *System_Dns_HostEntryPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Dns_HostEntry) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/alias with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_AliasPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, true)
	if ok {
		return convertSystem_Dns_HostEntry_AliasPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/alias with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_AliasPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/alias with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_AliasPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
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
		qv := convertSystem_Dns_HostEntry_AliasPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/alias with a ONCE subscription.
func (n *System_Dns_HostEntry_AliasPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/alias.
func (n *System_Dns_HostEntry_AliasPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/host-entries/host-entry/config/alias in the given batch object.
func (n *System_Dns_HostEntry_AliasPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/alias.
func (n *System_Dns_HostEntry_AliasPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/host-entries/host-entry/config/alias in the given batch object.
func (n *System_Dns_HostEntry_AliasPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/alias.
func (n *System_Dns_HostEntry_AliasPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/host-entries/host-entry/config/alias in the given batch object.
func (n *System_Dns_HostEntry_AliasPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Dns_HostEntry_AliasPath extracts the value of the leaf Alias from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_AliasPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Alias
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/hostname with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_HostnamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, true)
	if ok {
		return convertSystem_Dns_HostEntry_HostnamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/hostname with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_HostnamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/hostname with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_HostnamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_HostnamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/hostname with a ONCE subscription.
func (n *System_Dns_HostEntry_HostnamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/hostname.
func (n *System_Dns_HostEntry_HostnamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/host-entries/host-entry/config/hostname in the given batch object.
func (n *System_Dns_HostEntry_HostnamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/hostname.
func (n *System_Dns_HostEntry_HostnamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/host-entries/host-entry/config/hostname in the given batch object.
func (n *System_Dns_HostEntry_HostnamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/hostname.
func (n *System_Dns_HostEntry_HostnamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/host-entries/host-entry/config/hostname in the given batch object.
func (n *System_Dns_HostEntry_HostnamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Dns_HostEntry_HostnamePath extracts the value of the leaf Hostname from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Dns_HostEntry_HostnamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, true)
	if ok {
		return convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
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
		qv := convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address with a ONCE subscription.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address in the given batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address in the given batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/dns/host-entries/host-entry/config/ipv4-address in the given batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Dns_HostEntry_Ipv4AddressPath extracts the value of the leaf Ipv4Address from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_Ipv4AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Ipv4Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

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
