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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *KeychainPath) Lookup(t testing.TB) *oc.QualifiedKeychain {
	t.Helper()
	goStruct := &oc.Keychain{}
	md, ok := oc.Lookup(t, n, "Keychain", goStruct, false, true)
	if ok {
		return (&oc.QualifiedKeychain{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *KeychainPath) Get(t testing.TB) *oc.Keychain {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *KeychainPathAny) Lookup(t testing.TB) []*oc.QualifiedKeychain {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedKeychain
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedKeychain{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain with a ONCE subscription.
func (n *KeychainPathAny) Get(t testing.TB) []*oc.Keychain {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Keychain
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain.
func (n *KeychainPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain in the given batch object.
func (n *KeychainPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain.
func (n *KeychainPath) Replace(t testing.TB, val *oc.Keychain) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain in the given batch object.
func (n *KeychainPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain.
func (n *KeychainPath) Update(t testing.TB, val *oc.Keychain) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain in the given batch object.
func (n *KeychainPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_KeyPath) Lookup(t testing.TB) *oc.QualifiedKeychain_Key {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, false, true)
	if ok {
		return (&oc.QualifiedKeychain_Key{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_KeyPath) Get(t testing.TB) *oc.Keychain_Key {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_KeyPathAny) Lookup(t testing.TB) []*oc.QualifiedKeychain_Key {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedKeychain_Key
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedKeychain_Key{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key with a ONCE subscription.
func (n *Keychain_KeyPathAny) Get(t testing.TB) []*oc.Keychain_Key {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Keychain_Key
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key.
func (n *Keychain_KeyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key in the given batch object.
func (n *Keychain_KeyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key.
func (n *Keychain_KeyPath) Replace(t testing.TB, val *oc.Keychain_Key) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key in the given batch object.
func (n *Keychain_KeyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain_Key) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key.
func (n *Keychain_KeyPath) Update(t testing.TB, val *oc.Keychain_Key) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key in the given batch object.
func (n *Keychain_KeyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Keychain_Key) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_CryptoAlgorithmPath) Lookup(t testing.TB) *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, true, true)
	if ok {
		return convertKeychain_Key_CryptoAlgorithmPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_CryptoAlgorithmPath) Get(t testing.TB) oc.E_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Lookup(t testing.TB) []*oc.QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_KeychainTypes_CRYPTO_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_CryptoAlgorithmPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm with a ONCE subscription.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Get(t testing.TB) []oc.E_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_KeychainTypes_CRYPTO_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm.
func (n *Keychain_Key_CryptoAlgorithmPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm in the given batch object.
func (n *Keychain_Key_CryptoAlgorithmPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm.
func (n *Keychain_Key_CryptoAlgorithmPath) Replace(t testing.TB, val oc.E_KeychainTypes_CRYPTO_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm in the given batch object.
func (n *Keychain_Key_CryptoAlgorithmPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_KeychainTypes_CRYPTO_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm.
func (n *Keychain_Key_CryptoAlgorithmPath) Update(t testing.TB, val oc.E_KeychainTypes_CRYPTO_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-keychain/keychains/keychain/keys/key/config/crypto-algorithm in the given batch object.
func (n *Keychain_Key_CryptoAlgorithmPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_KeychainTypes_CRYPTO_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertKeychain_Key_CryptoAlgorithmPath extracts the value of the leaf CryptoAlgorithm from its parent oc.Keychain_Key
// and combines the update with an existing Metadata to return a *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE.
func convertKeychain_Key_CryptoAlgorithmPath(t testing.TB, md *genutil.Metadata, parent *oc.Keychain_Key) *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_KeychainTypes_CRYPTO_TYPE{
		Metadata: md,
	}
	val := parent.CryptoAlgorithm
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
