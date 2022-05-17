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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetime_EndTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, true, true)
	if ok {
		return convertKeychain_Key_SendLifetime_EndTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetime_EndTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_SendLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_SendLifetime_EndTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time with a ONCE subscription.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time.
func (n *Keychain_Key_SendLifetime_EndTimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time in the given batch object.
func (n *Keychain_Key_SendLifetime_EndTimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time.
func (n *Keychain_Key_SendLifetime_EndTimePath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time in the given batch object.
func (n *Keychain_Key_SendLifetime_EndTimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time.
func (n *Keychain_Key_SendLifetime_EndTimePath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/end-time in the given batch object.
func (n *Keychain_Key_SendLifetime_EndTimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_SendLifetime_EndTimePath extracts the value of the leaf EndTime from its parent oc.Keychain_Key_SendLifetime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertKeychain_Key_SendLifetime_EndTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key_SendLifetime) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, true, true)
	if ok {
		return convertKeychain_Key_SendLifetime_SendAndReceivePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSendAndReceive())
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_SendLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_SendLifetime_SendAndReceivePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive with a ONCE subscription.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive in the given batch object.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive in the given batch object.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/send-and-receive in the given batch object.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_SendLifetime_SendAndReceivePath extracts the value of the leaf SendAndReceive from its parent oc.Keychain_Key_SendLifetime
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertKeychain_Key_SendLifetime_SendAndReceivePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key_SendLifetime) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.SendAndReceive
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetime_StartTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, true, true)
	if ok {
		return convertKeychain_Key_SendLifetime_StartTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetime_StartTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_SendLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_SendLifetime_StartTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time with a ONCE subscription.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time.
func (n *Keychain_Key_SendLifetime_StartTimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time in the given batch object.
func (n *Keychain_Key_SendLifetime_StartTimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time.
func (n *Keychain_Key_SendLifetime_StartTimePath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time in the given batch object.
func (n *Keychain_Key_SendLifetime_StartTimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time.
func (n *Keychain_Key_SendLifetime_StartTimePath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/config/start-time in the given batch object.
func (n *Keychain_Key_SendLifetime_StartTimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_Key_SendLifetime_StartTimePath extracts the value of the leaf StartTime from its parent oc.Keychain_Key_SendLifetime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertKeychain_Key_SendLifetime_StartTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key_SendLifetime) *oc.QualifiedUint64 {
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
