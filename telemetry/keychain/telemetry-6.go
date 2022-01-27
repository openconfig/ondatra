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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Keychain{}
	md, ok := oc.Lookup(t, n, "Keychain", goStruct, true, false)
	if ok {
		return convertKeychain_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/state/name with a ONCE subscription.
func (n *Keychain_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Keychain{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain", gs, queryPath, true, false)
		return convertKeychain_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Keychain_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/state/name to the batch object.
func (n *Keychain_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Keychain_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/state/name to the batch object.
func (n *Keychain_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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

// Lookup fetches the value at /openconfig-keychain/keychains/keychain/state/tolerance with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Keychain_TolerancePath) Lookup(t testing.TB) *oc.QualifiedKeychain_Tolerance_Union {
	t.Helper()
	goStruct := &oc.Keychain{}
	md, ok := oc.Lookup(t, n, "Keychain", goStruct, true, false)
	if ok {
		return convertKeychain_TolerancePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-keychain/keychains/keychain/state/tolerance with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Keychain_TolerancePath) Get(t testing.TB) oc.Keychain_Tolerance_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-keychain/keychains/keychain/state/tolerance with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Keychain_TolerancePathAny) Lookup(t testing.TB) []*oc.QualifiedKeychain_Tolerance_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedKeychain_Tolerance_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Keychain{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Keychain", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertKeychain_TolerancePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-keychain/keychains/keychain/state/tolerance with a ONCE subscription.
func (n *Keychain_TolerancePathAny) Get(t testing.TB) []oc.Keychain_Tolerance_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Keychain_Tolerance_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/state/tolerance with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_TolerancePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Tolerance_Union {
	t.Helper()
	c := &oc.CollectionKeychain_Tolerance_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Tolerance_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Keychain_TolerancePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedKeychain_Tolerance_Union) bool) *oc.Keychain_Tolerance_UnionWatcher {
	t.Helper()
	w := &oc.Keychain_Tolerance_UnionWatcher{}
	gs := &oc.Keychain{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Keychain", gs, queryPath, true, false)
		return convertKeychain_TolerancePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedKeychain_Tolerance_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/state/tolerance with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_TolerancePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Tolerance_Union) bool) *oc.Keychain_Tolerance_UnionWatcher {
	t.Helper()
	return watch_Keychain_TolerancePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-keychain/keychains/keychain/state/tolerance with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Keychain_TolerancePath) Await(t testing.TB, timeout time.Duration, val oc.Keychain_Tolerance_Union) *oc.QualifiedKeychain_Tolerance_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedKeychain_Tolerance_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-keychain/keychains/keychain/state/tolerance failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-keychain/keychains/keychain/state/tolerance to the batch object.
func (n *Keychain_TolerancePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-keychain/keychains/keychain/state/tolerance with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Keychain_TolerancePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionKeychain_Tolerance_Union {
	t.Helper()
	c := &oc.CollectionKeychain_Tolerance_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedKeychain_Tolerance_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-keychain/keychains/keychain/state/tolerance with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Keychain_TolerancePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedKeychain_Tolerance_Union) bool) *oc.Keychain_Tolerance_UnionWatcher {
	t.Helper()
	return watch_Keychain_TolerancePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-keychain/keychains/keychain/state/tolerance to the batch object.
func (n *Keychain_TolerancePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
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
