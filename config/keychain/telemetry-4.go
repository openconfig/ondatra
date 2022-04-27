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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Keychain{}
	md, ok := oc.Lookup(t, n, "Keychain", goStruct, true, true)
	if ok {
		return convertKeychain_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/config/name with a ONCE subscription.
func (n *Keychain_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/config/name.
func (n *Keychain_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/config/name in the given batch object.
func (n *Keychain_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/config/name.
func (n *Keychain_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/config/name in the given batch object.
func (n *Keychain_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/config/name.
func (n *Keychain_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/config/name in the given batch object.
func (n *Keychain_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertKeychain_NamePath extracts the value of the leaf Name from its parent oc.Keychain
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertKeychain_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/config/tolerance with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_TolerancePath) Lookup(t testing.TB) *oc.QualifiedKeychain_Tolerance_Union {
	t.Helper()
	goStruct := &oc.Keychain{}
	md, ok := oc.Lookup(t, n, "Keychain", goStruct, true, true)
	if ok {
		return convertKeychain_TolerancePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/config/tolerance with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_TolerancePath) Get(t testing.TB) oc.Keychain_Tolerance_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/config/tolerance with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_TolerancePathAny) Lookup(t testing.TB) []*oc.QualifiedKeychain_Tolerance_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedKeychain_Tolerance_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_TolerancePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/config/tolerance with a ONCE subscription.
func (n *Keychain_TolerancePathAny) Get(t testing.TB) []oc.Keychain_Tolerance_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Keychain_Tolerance_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/config/tolerance.
func (n *Keychain_TolerancePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/config/tolerance in the given batch object.
func (n *Keychain_TolerancePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/config/tolerance.
func (n *Keychain_TolerancePath) Replace(t testing.TB, val oc.Keychain_Tolerance_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/config/tolerance in the given batch object.
func (n *Keychain_TolerancePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Keychain_Tolerance_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/config/tolerance.
func (n *Keychain_TolerancePath) Update(t testing.TB, val oc.Keychain_Tolerance_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/config/tolerance in the given batch object.
func (n *Keychain_TolerancePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Keychain_Tolerance_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertKeychain_TolerancePath extracts the value of the leaf Tolerance from its parent oc.Keychain
// and combines the update with an existing Metadata to return a *oc.QualifiedKeychain_Tolerance_Union.
func convertKeychain_TolerancePath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain) *oc.QualifiedKeychain_Tolerance_Union {
	t.Helper()
	qv := &oc.QualifiedKeychain_Tolerance_Union{
		Metadata: md,
	}
	val := parent.Tolerance
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
