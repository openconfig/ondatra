package system

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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Tacacs{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath extracts the value of the leaf SecretKeyHashed from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SecretKeyHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Tacacs{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath extracts the value of the leaf SecretKey from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SecretKey
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Tacacs{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_TimeoutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_TimeoutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout to the batch object.
func (n *System_Aaa_ServerGroup_Server_TimeoutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_TimeoutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_TimeoutPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout to the batch object.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_TimeoutPath extracts the value of the leaf Timeout from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_TypePath) Lookup(t testing.TB) *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/state/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_TypePath) Get(t testing.TB) oc.E_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a ONCE subscription.
func (n *System_Aaa_ServerGroup_TypePathAny) Get(t testing.TB) []oc.E_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AaaTypes_AAA_SERVER_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	c := &oc.CollectionE_AaaTypes_AAA_SERVER_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool) *oc.E_AaaTypes_AAA_SERVER_TYPEWatcher {
	t.Helper()
	w := &oc.E_AaaTypes_AAA_SERVER_TYPEWatcher{}
	gs := &oc.System_Aaa_ServerGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_TypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool) *oc.E_AaaTypes_AAA_SERVER_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_AaaTypes_AAA_SERVER_TYPE) *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/state/type to the batch object.
func (n *System_Aaa_ServerGroup_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	c := &oc.CollectionE_AaaTypes_AAA_SERVER_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_TypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool) *oc.E_AaaTypes_AAA_SERVER_TYPEWatcher {
	t.Helper()
	w := &oc.E_AaaTypes_AAA_SERVER_TYPEWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_TypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool) *oc.E_AaaTypes_AAA_SERVER_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_TypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/state/type to the batch object.
func (n *System_Aaa_ServerGroup_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_TypePath extracts the value of the leaf Type from its parent oc.System_Aaa_ServerGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE.
func convertSystem_Aaa_ServerGroup_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup) *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_AlarmPath) Lookup(t testing.TB) *oc.QualifiedSystem_Alarm {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Alarm{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_AlarmPath) Get(t testing.TB) *oc.System_Alarm {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_AlarmPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Alarm {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Alarm
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Alarm{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm with a ONCE subscription.
func (n *System_AlarmPathAny) Get(t testing.TB) []*oc.System_Alarm {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Alarm
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_AlarmPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Alarm {
	t.Helper()
	c := &oc.CollectionSystem_Alarm{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Alarm) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Alarm{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Alarm)))
		return false
	})
	return c
}

func watch_System_AlarmPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Alarm) bool) *oc.System_AlarmWatcher {
	t.Helper()
	w := &oc.System_AlarmWatcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Alarm{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Alarm)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_AlarmPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Alarm) bool) *oc.System_AlarmWatcher {
	t.Helper()
	return watch_System_AlarmPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_AlarmPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Alarm) *oc.QualifiedSystem_Alarm {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Alarm) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm to the batch object.
func (n *System_AlarmPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_AlarmPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Alarm {
	t.Helper()
	c := &oc.CollectionSystem_Alarm{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Alarm) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_AlarmPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Alarm) bool) *oc.System_AlarmWatcher {
	t.Helper()
	w := &oc.System_AlarmWatcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Alarm{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Alarm)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_AlarmPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Alarm) bool) *oc.System_AlarmWatcher {
	t.Helper()
	return watch_System_AlarmPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm to the batch object.
func (n *System_AlarmPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Alarm_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, true, false)
	if ok {
		return convertSystem_Alarm_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Alarm_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Alarm_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Alarm_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm/state/id with a ONCE subscription.
func (n *System_Alarm_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Alarm_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Alarm_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm/state/id to the batch object.
func (n *System_Alarm_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, true, false)
			qv := convertSystem_Alarm_IdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm/state/id to the batch object.
func (n *System_Alarm_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Alarm_IdPath extracts the value of the leaf Id from its parent oc.System_Alarm
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Alarm_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Alarm) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm/state/resource with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Alarm_ResourcePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, true, false)
	if ok {
		return convertSystem_Alarm_ResourcePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm/state/resource with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Alarm_ResourcePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm/state/resource with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Alarm_ResourcePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Alarm_ResourcePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm/state/resource with a ONCE subscription.
func (n *System_Alarm_ResourcePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/resource with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_ResourcePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_ResourcePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Alarm_ResourcePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_ResourcePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_ResourcePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm/state/resource with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Alarm_ResourcePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm/state/resource failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm/state/resource to the batch object.
func (n *System_Alarm_ResourcePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/resource with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_ResourcePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_ResourcePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, true, false)
			qv := convertSystem_Alarm_ResourcePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_ResourcePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_ResourcePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm/state/resource to the batch object.
func (n *System_Alarm_ResourcePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Alarm_ResourcePath extracts the value of the leaf Resource from its parent oc.System_Alarm
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Alarm_ResourcePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Alarm) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Resource
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm/state/severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Alarm_SeverityPath) Lookup(t testing.TB) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, true, false)
	if ok {
		return convertSystem_Alarm_SeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm/state/severity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Alarm_SeverityPath) Get(t testing.TB) oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm/state/severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Alarm_SeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Alarm_SeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm/state/severity with a ONCE subscription.
func (n *System_Alarm_SeverityPathAny) Get(t testing.TB) []oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_SeverityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c := &oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_SeverityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	w := &oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Alarm_SeverityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_SeverityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_System_Alarm_SeverityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm/state/severity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Alarm_SeverityPath) Await(t testing.TB, timeout time.Duration, val oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm/state/severity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm/state/severity to the batch object.
func (n *System_Alarm_SeverityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_SeverityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c := &oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_SeverityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	w := &oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, true, false)
			qv := convertSystem_Alarm_SeverityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_SeverityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_System_Alarm_SeverityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm/state/severity to the batch object.
func (n *System_Alarm_SeverityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Alarm_SeverityPath extracts the value of the leaf Severity from its parent oc.System_Alarm
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY.
func convertSystem_Alarm_SeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Alarm) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	qv := &oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{
		Metadata: md,
	}
	val := parent.Severity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm/state/text with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Alarm_TextPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, true, false)
	if ok {
		return convertSystem_Alarm_TextPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm/state/text with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Alarm_TextPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm/state/text with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Alarm_TextPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Alarm_TextPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm/state/text with a ONCE subscription.
func (n *System_Alarm_TextPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/text with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_TextPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_TextPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Alarm_TextPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/text with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TextPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_TextPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm/state/text with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Alarm_TextPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm/state/text failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm/state/text to the batch object.
func (n *System_Alarm_TextPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/text with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_TextPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_TextPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, true, false)
			qv := convertSystem_Alarm_TextPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/text with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TextPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_TextPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm/state/text to the batch object.
func (n *System_Alarm_TextPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Alarm_TextPath extracts the value of the leaf Text from its parent oc.System_Alarm
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Alarm_TextPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Alarm) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Text
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm/state/time-created with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Alarm_TimeCreatedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, true, false)
	if ok {
		return convertSystem_Alarm_TimeCreatedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm/state/time-created with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Alarm_TimeCreatedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm/state/time-created with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Alarm_TimeCreatedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Alarm_TimeCreatedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm/state/time-created with a ONCE subscription.
func (n *System_Alarm_TimeCreatedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/time-created with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_TimeCreatedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_TimeCreatedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Alarm_TimeCreatedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/time-created with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TimeCreatedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Alarm_TimeCreatedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm/state/time-created with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Alarm_TimeCreatedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm/state/time-created failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm/state/time-created to the batch object.
func (n *System_Alarm_TimeCreatedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/time-created with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_TimeCreatedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_TimeCreatedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, true, false)
			qv := convertSystem_Alarm_TimeCreatedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/time-created with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TimeCreatedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Alarm_TimeCreatedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm/state/time-created to the batch object.
func (n *System_Alarm_TimeCreatedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Alarm_TimeCreatedPath extracts the value of the leaf TimeCreated from its parent oc.System_Alarm
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Alarm_TimeCreatedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Alarm) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TimeCreated
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/alarms/alarm/state/type-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Alarm_TypeIdPath) Lookup(t testing.TB) *oc.QualifiedSystem_Alarm_TypeId_Union {
	t.Helper()
	goStruct := &oc.System_Alarm{}
	md, ok := oc.Lookup(t, n, "System_Alarm", goStruct, true, false)
	if ok {
		return convertSystem_Alarm_TypeIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/alarms/alarm/state/type-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Alarm_TypeIdPath) Get(t testing.TB) oc.System_Alarm_TypeId_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/alarms/alarm/state/type-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Alarm_TypeIdPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Alarm_TypeId_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Alarm_TypeId_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Alarm{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Alarm", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Alarm_TypeIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/alarms/alarm/state/type-id with a ONCE subscription.
func (n *System_Alarm_TypeIdPathAny) Get(t testing.TB) []oc.System_Alarm_TypeId_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_Alarm_TypeId_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/type-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_TypeIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Alarm_TypeId_Union {
	t.Helper()
	c := &oc.CollectionSystem_Alarm_TypeId_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Alarm_TypeId_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_TypeIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Alarm_TypeId_Union) bool) *oc.System_Alarm_TypeId_UnionWatcher {
	t.Helper()
	w := &oc.System_Alarm_TypeId_UnionWatcher{}
	gs := &oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Alarm_TypeIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Alarm_TypeId_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/type-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TypeIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Alarm_TypeId_Union) bool) *oc.System_Alarm_TypeId_UnionWatcher {
	t.Helper()
	return watch_System_Alarm_TypeIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/alarms/alarm/state/type-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Alarm_TypeIdPath) Await(t testing.TB, timeout time.Duration, val oc.System_Alarm_TypeId_Union) *oc.QualifiedSystem_Alarm_TypeId_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Alarm_TypeId_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/alarms/alarm/state/type-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/alarms/alarm/state/type-id to the batch object.
func (n *System_Alarm_TypeIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/alarms/alarm/state/type-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Alarm_TypeIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Alarm_TypeId_Union {
	t.Helper()
	c := &oc.CollectionSystem_Alarm_TypeId_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Alarm_TypeId_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Alarm_TypeIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Alarm_TypeId_Union) bool) *oc.System_Alarm_TypeId_UnionWatcher {
	t.Helper()
	w := &oc.System_Alarm_TypeId_UnionWatcher{}
	structs := map[string]*oc.System_Alarm{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Alarm{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Alarm", structs[pre], queryPath, true, false)
			qv := convertSystem_Alarm_TypeIdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Alarm_TypeId_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/type-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TypeIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Alarm_TypeId_Union) bool) *oc.System_Alarm_TypeId_UnionWatcher {
	t.Helper()
	return watch_System_Alarm_TypeIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/alarms/alarm/state/type-id to the batch object.
func (n *System_Alarm_TypeIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Alarm_TypeIdPath extracts the value of the leaf TypeId from its parent oc.System_Alarm
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Alarm_TypeId_Union.
func convertSystem_Alarm_TypeIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Alarm) *oc.QualifiedSystem_Alarm_TypeId_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_Alarm_TypeId_Union{
		Metadata: md,
	}
	val := parent.TypeId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/state/boot-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_BootTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, false)
	if ok {
		return convertSystem_BootTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/state/boot-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_BootTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/state/boot-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_BootTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_BootTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/state/boot-time with a ONCE subscription.
func (n *System_BootTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/boot-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_BootTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_BootTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_BootTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/boot-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_BootTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_BootTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/state/boot-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_BootTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/state/boot-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/state/boot-time to the batch object.
func (n *System_BootTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/boot-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_BootTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_BootTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System", structs[pre], queryPath, true, false)
			qv := convertSystem_BootTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/boot-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_BootTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_BootTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/state/boot-time to the batch object.
func (n *System_BootTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_BootTimePath extracts the value of the leaf BootTime from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_BootTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.BootTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/clock with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_ClockPath) Lookup(t testing.TB) *oc.QualifiedSystem_Clock {
	t.Helper()
	goStruct := &oc.System_Clock{}
	md, ok := oc.Lookup(t, n, "System_Clock", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Clock{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/clock with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_ClockPath) Get(t testing.TB) *oc.System_Clock {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/clock with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_ClockPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Clock {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Clock
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Clock{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Clock", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Clock{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/clock with a ONCE subscription.
func (n *System_ClockPathAny) Get(t testing.TB) []*oc.System_Clock {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Clock
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/clock with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_ClockPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Clock {
	t.Helper()
	c := &oc.CollectionSystem_Clock{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Clock) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Clock{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Clock)))
		return false
	})
	return c
}

func watch_System_ClockPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Clock) bool) *oc.System_ClockWatcher {
	t.Helper()
	w := &oc.System_ClockWatcher{}
	gs := &oc.System_Clock{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Clock", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Clock{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Clock)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/clock with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_ClockPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Clock) bool) *oc.System_ClockWatcher {
	t.Helper()
	return watch_System_ClockPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/clock with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_ClockPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Clock) *oc.QualifiedSystem_Clock {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Clock) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/clock failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/clock to the batch object.
func (n *System_ClockPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/clock with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_ClockPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Clock {
	t.Helper()
	c := &oc.CollectionSystem_Clock{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Clock) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_ClockPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Clock) bool) *oc.System_ClockWatcher {
	t.Helper()
	w := &oc.System_ClockWatcher{}
	structs := map[string]*oc.System_Clock{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Clock{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Clock", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Clock{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Clock)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/clock with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_ClockPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Clock) bool) *oc.System_ClockWatcher {
	t.Helper()
	return watch_System_ClockPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/clock to the batch object.
func (n *System_ClockPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/clock/state/timezone-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Clock_TimezoneNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Clock{}
	md, ok := oc.Lookup(t, n, "System_Clock", goStruct, true, false)
	if ok {
		return convertSystem_Clock_TimezoneNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/clock/state/timezone-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Clock_TimezoneNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/clock/state/timezone-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Clock_TimezoneNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Clock{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Clock", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Clock_TimezoneNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/clock/state/timezone-name with a ONCE subscription.
func (n *System_Clock_TimezoneNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/clock/state/timezone-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Clock_TimezoneNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Clock_TimezoneNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Clock{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Clock", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Clock_TimezoneNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/clock/state/timezone-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Clock_TimezoneNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Clock_TimezoneNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/clock/state/timezone-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Clock_TimezoneNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/clock/state/timezone-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/clock/state/timezone-name to the batch object.
func (n *System_Clock_TimezoneNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/clock/state/timezone-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Clock_TimezoneNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Clock_TimezoneNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Clock{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Clock{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Clock", structs[pre], queryPath, true, false)
			qv := convertSystem_Clock_TimezoneNamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/clock/state/timezone-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Clock_TimezoneNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Clock_TimezoneNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/clock/state/timezone-name to the batch object.
func (n *System_Clock_TimezoneNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Clock_TimezoneNamePath extracts the value of the leaf TimezoneName from its parent oc.System_Clock
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Clock_TimezoneNamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Clock) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TimezoneName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_CpuPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu {
	t.Helper()
	goStruct := &oc.System_Cpu{}
	md, ok := oc.Lookup(t, n, "System_Cpu", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_CpuPath) Get(t testing.TB) *oc.System_Cpu {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_CpuPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu with a ONCE subscription.
func (n *System_CpuPathAny) Get(t testing.TB) []*oc.System_Cpu {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_CpuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu {
	t.Helper()
	c := &oc.CollectionSystem_Cpu{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu)))
		return false
	})
	return c
}

func watch_System_CpuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu) bool) *oc.System_CpuWatcher {
	t.Helper()
	w := &oc.System_CpuWatcher{}
	gs := &oc.System_Cpu{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Cpu{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_CpuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu) bool) *oc.System_CpuWatcher {
	t.Helper()
	return watch_System_CpuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_CpuPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu) *oc.QualifiedSystem_Cpu {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu to the batch object.
func (n *System_CpuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_CpuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu {
	t.Helper()
	c := &oc.CollectionSystem_Cpu{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_CpuPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu) bool) *oc.System_CpuWatcher {
	t.Helper()
	w := &oc.System_CpuWatcher{}
	structs := map[string]*oc.System_Cpu{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Cpu{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_CpuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu) bool) *oc.System_CpuWatcher {
	t.Helper()
	return watch_System_CpuPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu to the batch object.
func (n *System_CpuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterruptPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_HardwareInterrupt {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_HardwareInterrupt{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterruptPath) Get(t testing.TB) *oc.System_Cpu_HardwareInterrupt {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterruptPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_HardwareInterrupt {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_HardwareInterrupt
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_HardwareInterrupt{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a ONCE subscription.
func (n *System_Cpu_HardwareInterruptPathAny) Get(t testing.TB) []*oc.System_Cpu_HardwareInterrupt {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_HardwareInterrupt
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterruptPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_HardwareInterrupt {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_HardwareInterrupt{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_HardwareInterrupt{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_HardwareInterrupt)))
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterruptPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool) *oc.System_Cpu_HardwareInterruptWatcher {
	t.Helper()
	w := &oc.System_Cpu_HardwareInterruptWatcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Cpu_HardwareInterrupt{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_HardwareInterrupt)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterruptPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool) *oc.System_Cpu_HardwareInterruptWatcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterruptPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterruptPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedSystem_Cpu_HardwareInterrupt {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt to the batch object.
func (n *System_Cpu_HardwareInterruptPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterruptPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_HardwareInterrupt {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_HardwareInterrupt{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterruptPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool) *oc.System_Cpu_HardwareInterruptWatcher {
	t.Helper()
	w := &oc.System_Cpu_HardwareInterruptWatcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Cpu_HardwareInterrupt{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_HardwareInterrupt)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterruptPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_HardwareInterrupt) bool) *oc.System_Cpu_HardwareInterruptWatcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterruptPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt to the batch object.
func (n *System_Cpu_HardwareInterruptPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg to the batch object.
func (n *System_Cpu_HardwareInterrupt_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_AvgPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/avg to the batch object.
func (n *System_Cpu_HardwareInterrupt_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_HardwareInterrupt_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant to the batch object.
func (n *System_Cpu_HardwareInterrupt_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_InstantPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/instant to the batch object.
func (n *System_Cpu_HardwareInterrupt_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_HardwareInterrupt_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval to the batch object.
func (n *System_Cpu_HardwareInterrupt_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/interval to the batch object.
func (n *System_Cpu_HardwareInterrupt_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_HardwareInterrupt_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max to the batch object.
func (n *System_Cpu_HardwareInterrupt_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_MaxPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max to the batch object.
func (n *System_Cpu_HardwareInterrupt_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_HardwareInterrupt_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time to the batch object.
func (n *System_Cpu_HardwareInterrupt_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/max-time to the batch object.
func (n *System_Cpu_HardwareInterrupt_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_HardwareInterrupt_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_HardwareInterrupt_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_HardwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_HardwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_HardwareInterrupt_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_IdlePath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Idle {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_IdlePath) Get(t testing.TB) *oc.System_Cpu_Idle {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_IdlePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Idle {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Idle
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription.
func (n *System_Cpu_IdlePathAny) Get(t testing.TB) []*oc.System_Cpu_Idle {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_Idle
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_IdlePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Idle {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Idle{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Idle) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_Idle)))
		return false
	})
	return c
}

func watch_System_Cpu_IdlePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	w := &oc.System_Cpu_IdleWatcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Idle)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_IdlePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	return watch_System_Cpu_IdlePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_IdlePath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_Idle) *oc.QualifiedSystem_Cpu_Idle {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Idle) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle to the batch object.
func (n *System_Cpu_IdlePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_IdlePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Idle {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Idle{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Idle) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_IdlePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	w := &oc.System_Cpu_IdleWatcher{}
	structs := map[string]*oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Idle{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Idle", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Cpu_Idle{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Idle)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_IdlePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	return watch_System_Cpu_IdlePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle to the batch object.
func (n *System_Cpu_IdlePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
