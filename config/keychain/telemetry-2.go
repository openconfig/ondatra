package keychain

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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_ReceiveLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_ReceiveLifetime", goStruct, true, true)
	if ok {
		return convertKeychain_Key_ReceiveLifetime_StartTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_ReceiveLifetime_StartTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_ReceiveLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_ReceiveLifetime", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_ReceiveLifetime_StartTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time with a ONCE subscription.
func (n *Keychain_Key_ReceiveLifetime_StartTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time in the given batch object.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time in the given batch object.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/start-time in the given batch object.
func (n *Keychain_Key_ReceiveLifetime_StartTimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_ReceiveLifetime_StartTimePath extracts the value of the leaf StartTime from its parent oc.Keychain_Key_ReceiveLifetime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertKeychain_Key_ReceiveLifetime_StartTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key_ReceiveLifetime) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.StartTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SecretKeyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, true, true)
	if ok {
		return convertKeychain_Key_SecretKeyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SecretKeyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SecretKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_SecretKeyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key with a ONCE subscription.
func (n *Keychain_Key_SecretKeyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key.
func (n *Keychain_Key_SecretKeyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key in the given batch object.
func (n *Keychain_Key_SecretKeyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key.
func (n *Keychain_Key_SecretKeyPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key in the given batch object.
func (n *Keychain_Key_SecretKeyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key.
func (n *Keychain_Key_SecretKeyPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/config/secret-key in the given batch object.
func (n *Keychain_Key_SecretKeyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_SecretKeyPath extracts the value of the leaf SecretKey from its parent oc.Keychain_Key
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertKeychain_Key_SecretKeyPath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetimePath) Lookup(t testing.TB) *oc.QualifiedKeychain_Key_SendLifetime {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, false, true)
	if ok {
		return (&oc.QualifiedKeychain_Key_SendLifetime{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetimePath) Get(t testing.TB) *oc.Keychain_Key_SendLifetime {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedKeychain_Key_SendLifetime {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedKeychain_Key_SendLifetime
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_SendLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedKeychain_Key_SendLifetime{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a ONCE subscription.
func (n *Keychain_Key_SendLifetimePathAny) Get(t testing.TB) []*oc.Keychain_Key_SendLifetime {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Keychain_Key_SendLifetime
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime.
func (n *Keychain_Key_SendLifetimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime in the given batch object.
func (n *Keychain_Key_SendLifetimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime.
func (n *Keychain_Key_SendLifetimePath) Replace(t testing.TB, val *oc.Keychain_Key_SendLifetime) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime in the given batch object.
func (n *Keychain_Key_SendLifetimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain_Key_SendLifetime) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime.
func (n *Keychain_Key_SendLifetimePath) Update(t testing.TB, val *oc.Keychain_Key_SendLifetime) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime in the given batch object.
func (n *Keychain_Key_SendLifetimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain_Key_SendLifetime) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
