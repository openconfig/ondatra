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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *KeychainPath) Lookup(t testing.TB) *oc.QualifiedKeychain {
	t.Helper()
	goStruct := &oc.Keychain{}
	md, ok := oc.Lookup(t, n, "Keychain", goStruct, false, false)
	if ok {
		return (&oc.QualifiedKeychain{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain", goStruct, queryPath, false, false)
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

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *KeychainPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain {
	t.Helper()
	c := &oc.CollectionKeychain{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedKeychain{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Keychain)))
		return false
	})
	return c
}

func watch_KeychainPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain) bool) *oc.KeychainWatcher {
	t.Helper()
	w := &oc.KeychainWatcher{}
	gs := &oc.Keychain{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain", gs, queryPath, false, false)
		qv := (&oc.QualifiedKeychain{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *KeychainPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain) bool) *oc.KeychainWatcher {
	t.Helper()
	return watch_KeychainPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *KeychainPath) Await(t testing.TB, timeout time.Duration, val *oc.Keychain) *oc.QualifiedKeychain {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedKeychain) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain to the batch object.
func (n *KeychainPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *KeychainPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain {
	t.Helper()
	c := &oc.CollectionKeychain{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_KeychainPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain) bool) *oc.KeychainWatcher {
	t.Helper()
	w := &oc.KeychainWatcher{}
	structs := map[string]*oc.Keychain{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedKeychain{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *KeychainPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain) bool) *oc.KeychainWatcher {
	t.Helper()
	return watch_KeychainPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain to the batch object.
func (n *KeychainPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_KeyPath) Lookup(t testing.TB) *oc.QualifiedKeychain_Key {
	t.Helper()
	goStruct := &oc.Keychain_Key{}
	md, ok := oc.Lookup(t, n, "Keychain_Key", goStruct, false, false)
	if ok {
		return (&oc.QualifiedKeychain_Key{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key", goStruct, queryPath, false, false)
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

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_KeyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Key {
	t.Helper()
	c := &oc.CollectionKeychain_Key{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Key) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedKeychain_Key{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Keychain_Key)))
		return false
	})
	return c
}

func watch_Keychain_KeyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Key) bool) *oc.Keychain_KeyWatcher {
	t.Helper()
	w := &oc.Keychain_KeyWatcher{}
	gs := &oc.Keychain_Key{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key", gs, queryPath, false, false)
		qv := (&oc.QualifiedKeychain_Key{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Key)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_KeyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Key) bool) *oc.Keychain_KeyWatcher {
	t.Helper()
	return watch_Keychain_KeyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_KeyPath) Await(t testing.TB, timeout time.Duration, val *oc.Keychain_Key) *oc.QualifiedKeychain_Key {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedKeychain_Key) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key to the batch object.
func (n *Keychain_KeyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_KeyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Key {
	t.Helper()
	c := &oc.CollectionKeychain_Key{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Key) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_KeyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Key) bool) *oc.Keychain_KeyWatcher {
	t.Helper()
	w := &oc.Keychain_KeyWatcher{}
	structs := map[string]*oc.Keychain_Key{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedKeychain_Key{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Key)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_KeyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Key) bool) *oc.Keychain_KeyWatcher {
	t.Helper()
	return watch_Keychain_KeyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key to the batch object.
func (n *Keychain_KeyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
