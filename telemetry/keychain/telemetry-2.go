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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_ReceiveLifetimePath) Lookup(t testing.TB) *oc.QualifiedKeychain_Key_ReceiveLifetime {
	t.Helper()
	goStruct := &oc.Keychain_Key_ReceiveLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_ReceiveLifetime", goStruct, false, false)
	if ok {
		return (&oc.QualifiedKeychain_Key_ReceiveLifetime{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_ReceiveLifetime", goStruct, queryPath, false, false)
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

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_ReceiveLifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Key_ReceiveLifetime {
	t.Helper()
	c := &oc.CollectionKeychain_Key_ReceiveLifetime{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Key_ReceiveLifetime) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedKeychain_Key_ReceiveLifetime{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Keychain_Key_ReceiveLifetime)))
		return false
	})
	return c
}

func watch_Keychain_Key_ReceiveLifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Key_ReceiveLifetime) bool) *oc.Keychain_Key_ReceiveLifetimeWatcher {
	t.Helper()
	w := &oc.Keychain_Key_ReceiveLifetimeWatcher{}
	gs := &oc.Keychain_Key_ReceiveLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key_ReceiveLifetime", gs, queryPath, false, false)
		qv := (&oc.QualifiedKeychain_Key_ReceiveLifetime{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Key_ReceiveLifetime)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_ReceiveLifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Key_ReceiveLifetime) bool) *oc.Keychain_Key_ReceiveLifetimeWatcher {
	t.Helper()
	return watch_Keychain_Key_ReceiveLifetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_ReceiveLifetimePath) Await(t testing.TB, timeout time.Duration, val *oc.Keychain_Key_ReceiveLifetime) *oc.QualifiedKeychain_Key_ReceiveLifetime {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedKeychain_Key_ReceiveLifetime) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime to the batch object.
func (n *Keychain_Key_ReceiveLifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_ReceiveLifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Key_ReceiveLifetime {
	t.Helper()
	c := &oc.CollectionKeychain_Key_ReceiveLifetime{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Key_ReceiveLifetime) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_ReceiveLifetimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Key_ReceiveLifetime) bool) *oc.Keychain_Key_ReceiveLifetimeWatcher {
	t.Helper()
	w := &oc.Keychain_Key_ReceiveLifetimeWatcher{}
	structs := map[string]*oc.Keychain_Key_ReceiveLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain_Key_ReceiveLifetime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key_ReceiveLifetime", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedKeychain_Key_ReceiveLifetime{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Key_ReceiveLifetime)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_ReceiveLifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Key_ReceiveLifetime) bool) *oc.Keychain_Key_ReceiveLifetimeWatcher {
	t.Helper()
	return watch_Keychain_Key_ReceiveLifetimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime to the batch object.
func (n *Keychain_Key_ReceiveLifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_ReceiveLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_ReceiveLifetime", goStruct, true, false)
	if ok {
		return convertKeychain_Key_ReceiveLifetime_EndTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_ReceiveLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_ReceiveLifetime", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_ReceiveLifetime_EndTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a ONCE subscription.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_ReceiveLifetime_EndTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Keychain_Key_ReceiveLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key_ReceiveLifetime", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertKeychain_Key_ReceiveLifetime_EndTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_ReceiveLifetime_EndTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time to the batch object.
func (n *Keychain_Key_ReceiveLifetime_EndTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_ReceiveLifetime_EndTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Keychain_Key_ReceiveLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain_Key_ReceiveLifetime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key_ReceiveLifetime", structs[pre], queryPath, true, false)
			qv := convertKeychain_Key_ReceiveLifetime_EndTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_ReceiveLifetime_EndTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/receive-lifetime/state/end-time to the batch object.
func (n *Keychain_Key_ReceiveLifetime_EndTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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
