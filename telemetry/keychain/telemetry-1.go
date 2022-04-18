package keychain

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_CryptoAlgorithmPath) Lookup(t testing.TB) *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, true, false)
	if ok {
		return convertKeychain_Key_CryptoAlgorithmPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_CryptoAlgorithmPath) Get(t testing.TB) oc.E_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Lookup(t testing.TB) []*oc.QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_KeychainTypes_CRYPTO_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_CryptoAlgorithmPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a ONCE subscription.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Get(t testing.TB) []oc.E_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_KeychainTypes_CRYPTO_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_CryptoAlgorithmPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	c := &oc.CollectionE_KeychainTypes_CRYPTO_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_CryptoAlgorithmPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool) *oc.E_KeychainTypes_CRYPTO_TYPEWatcher {
	t.Helper()
	w := &oc.E_KeychainTypes_CRYPTO_TYPEWatcher{}
	gs := &oc.Keychain_Key{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertKeychain_Key_CryptoAlgorithmPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_KeychainTypes_CRYPTO_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_CryptoAlgorithmPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool) *oc.E_KeychainTypes_CRYPTO_TYPEWatcher {
	t.Helper()
	return watch_Keychain_Key_CryptoAlgorithmPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_CryptoAlgorithmPath) Await(t testing.TB, timeout time.Duration, val oc.E_KeychainTypes_CRYPTO_TYPE) *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm to the batch object.
func (n *Keychain_Key_CryptoAlgorithmPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	c := &oc.CollectionE_KeychainTypes_CRYPTO_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_CryptoAlgorithmPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool) *oc.E_KeychainTypes_CRYPTO_TYPEWatcher {
	t.Helper()
	w := &oc.E_KeychainTypes_CRYPTO_TYPEWatcher{}
	structs := map[string]*oc.Keychain_Key{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain_Key{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key", structs[pre], queryPath, true, false)
			qv := convertKeychain_Key_CryptoAlgorithmPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_KeychainTypes_CRYPTO_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_KeychainTypes_CRYPTO_TYPE) bool) *oc.E_KeychainTypes_CRYPTO_TYPEWatcher {
	t.Helper()
	return watch_Keychain_Key_CryptoAlgorithmPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm to the batch object.
func (n *Keychain_Key_CryptoAlgorithmPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_KeyIdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, true, false)
	if ok {
		return convertKeychain_Key_KeyIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_KeyIdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_KeyIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_KeyIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a ONCE subscription.
func (n *Keychain_Key_KeyIdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_KeyIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_KeyIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Keychain_Key{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertKeychain_Key_KeyIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_KeyIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_KeyIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_KeyIdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/state/key-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/state/key-id to the batch object.
func (n *Keychain_Key_KeyIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_KeyIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_KeyIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Keychain_Key{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain_Key{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key", structs[pre], queryPath, true, false)
			qv := convertKeychain_Key_KeyIdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/state/key-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_KeyIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_KeyIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/state/key-id to the batch object.
func (n *Keychain_Key_KeyIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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
