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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, true, false)
	if ok {
		return convertKeychain_Key_SendLifetime_SendAndReceivePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSendAndReceive())
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain_Key_SendLifetime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain_Key_SendLifetime", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_Key_SendLifetime_SendAndReceivePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a ONCE subscription.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetime_SendAndReceivePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Keychain_Key_SendLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key_SendLifetime", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertKeychain_Key_SendLifetime_SendAndReceivePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetime_SendAndReceivePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive to the batch object.
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetime_SendAndReceivePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
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
			qv := convertKeychain_Key_SendLifetime_SendAndReceivePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetime_SendAndReceivePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive to the batch object.
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_Key_SendLifetime_StartTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Keychain_Key_SendLifetime{}
	md, ok := oc.Lookup(t, n, "Keychain_Key_SendLifetime", goStruct, true, false)
	if ok {
		return convertKeychain_Key_SendLifetime_StartTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_Key_SendLifetime_StartTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertKeychain_Key_SendLifetime_StartTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a ONCE subscription.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetime_StartTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetime_StartTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Keychain_Key_SendLifetime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain_Key_SendLifetime", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertKeychain_Key_SendLifetime_StartTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetime_StartTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetime_StartTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_Key_SendLifetime_StartTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time to the batch object.
func (n *Keychain_Key_SendLifetime_StartTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_Key_SendLifetime_StartTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertKeychain_Key_SendLifetime_StartTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Keychain_Key_SendLifetime_StartTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/start-time to the batch object.
func (n *Keychain_Key_SendLifetime_StartTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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
