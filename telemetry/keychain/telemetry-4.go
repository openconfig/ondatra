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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetimePath) Lookup(t testing.TB) *oc.QualifiedKeychain_Key_SendLifetime {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, false, false)
	if ok {
		return (&oc.QualifiedKeychain_Key_SendLifetime{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
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
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, false, false)
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

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Key_SendLifetime {
	t.Helper()
	c := &oc.CollectionKeychain_Key_SendLifetime{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Key_SendLifetime) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedKeychain_Key_SendLifetime{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Keychain_Key_SendLifetime)))
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Key_SendLifetime) bool) *oc.Keychain_Key_SendLifetimeWatcher {
	t.Helper()
	w := &oc.Keychain_Key_SendLifetimeWatcher{}
	gs := &oc.Keychain_Key_SendLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key_SendLifetime", gs, queryPath, false, false)
		qv := (&oc.QualifiedKeychain_Key_SendLifetime{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Key_SendLifetime)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Key_SendLifetime) bool) *oc.Keychain_Key_SendLifetimeWatcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_SendLifetimePath) Await(t testing.TB, timeout time.Duration, val *oc.Keychain_Key_SendLifetime) *oc.QualifiedKeychain_Key_SendLifetime {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedKeychain_Key_SendLifetime) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime to the batch object.
func (n *Keychain_Key_SendLifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Key_SendLifetime {
	t.Helper()
	c := &oc.CollectionKeychain_Key_SendLifetime{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Key_SendLifetime) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Key_SendLifetime) bool) *oc.Keychain_Key_SendLifetimeWatcher {
	t.Helper()
	w := &oc.Keychain_Key_SendLifetimeWatcher{}
	structs := map[string]*oc.Keychain_Key_SendLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain_Key_SendLifetime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key_SendLifetime", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedKeychain_Key_SendLifetime{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Key_SendLifetime)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Key_SendLifetime) bool) *oc.Keychain_Key_SendLifetimeWatcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime to the batch object.
func (n *Keychain_Key_SendLifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetime_EndTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, true, false)
	if ok {
		return convertKeychain_Key_SendLifetime_EndTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetime_EndTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_SendLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_SendLifetime_EndTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a ONCE subscription.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetime_EndTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetime_EndTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Keychain_Key_SendLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key_SendLifetime", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertKeychain_Key_SendLifetime_EndTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetime_EndTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetime_EndTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_SendLifetime_EndTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time to the batch object.
func (n *Keychain_Key_SendLifetime_EndTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetime_EndTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Keychain_Key_SendLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Keychain_Key_SendLifetime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Keychain_Key_SendLifetime", structs[pre], queryPath, true, false)
			qv := convertKeychain_Key_SendLifetime_EndTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetime_EndTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time to the batch object.
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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
