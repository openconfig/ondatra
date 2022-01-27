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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/config/key-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_KeyIdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, true, true)
	if ok {
		return convertKeychain_Key_KeyIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/config/key-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_KeyIdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/config/key-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_KeyIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_KeyIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/config/key-id with a ONCE subscription.
func (n *Keychain_Key_KeyIdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/key-id.
func (n *Keychain_Key_KeyIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/config/key-id in the given batch object.
func (n *Keychain_Key_KeyIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/key-id.
func (n *Keychain_Key_KeyIdPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/config/key-id in the given batch object.
func (n *Keychain_Key_KeyIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/key-id.
func (n *Keychain_Key_KeyIdPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/config/key-id in the given batch object.
func (n *Keychain_Key_KeyIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_KeyIdPath extracts the value of the leaf KeyId from its parent oc.Keychain_Key
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertKeychain_Key_KeyIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.KeyId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_ReceiveLifetimePath) Lookup(t testing.TB) *oc.QualifiedKeychain_Key_ReceiveLifetime {
	t.Helper()
	goStruct := &oc.Keychain_Key_ReceiveLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_ReceiveLifetime", goStruct, false, true)
	if ok {
		return (&oc.QualifiedKeychain_Key_ReceiveLifetime{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_ReceiveLifetimePath) Get(t testing.TB) *oc.Keychain_Key_ReceiveLifetime {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_ReceiveLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedKeychain_Key_ReceiveLifetime {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedKeychain_Key_ReceiveLifetime
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_ReceiveLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_ReceiveLifetime", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedKeychain_Key_ReceiveLifetime{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a ONCE subscription.
func (n *Keychain_Key_ReceiveLifetimePathAny) Get(t testing.TB) []*oc.Keychain_Key_ReceiveLifetime {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Keychain_Key_ReceiveLifetime
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime.
func (n *Keychain_Key_ReceiveLifetimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime in the given batch object.
func (n *Keychain_Key_ReceiveLifetimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime.
func (n *Keychain_Key_ReceiveLifetimePath) Replace(t testing.TB, val *oc.Keychain_Key_ReceiveLifetime) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime in the given batch object.
func (n *Keychain_Key_ReceiveLifetimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain_Key_ReceiveLifetime) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime.
func (n *Keychain_Key_ReceiveLifetimePath) Update(t testing.TB, val *oc.Keychain_Key_ReceiveLifetime) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime in the given batch object.
func (n *Keychain_Key_ReceiveLifetimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain_Key_ReceiveLifetime) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_ReceiveLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_ReceiveLifetime", goStruct, true, true)
	if ok {
		return convertKeychain_Key_ReceiveLifetime_EndTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertKeychain_Key_ReceiveLifetime_EndTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time with a ONCE subscription.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time in the given batch object.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time in the given batch object.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/config/end-time in the given batch object.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_ReceiveLifetime_EndTimePath extracts the value of the leaf EndTime from its parent oc.Keychain_Key_ReceiveLifetime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertKeychain_Key_ReceiveLifetime_EndTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key_ReceiveLifetime) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.EndTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
