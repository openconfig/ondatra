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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup_Server_Tacacs
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server_Tacacs{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_ServerGroup_Server_Tacacs)))
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_TacacsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) bool) *oc.System_Aaa_ServerGroup_Server_TacacsWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_Server_TacacsWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) bool) *oc.System_Aaa_ServerGroup_Server_TacacsWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_TacacsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs to the batch object.
func (n *System_Aaa_ServerGroup_Server_TacacsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server_Tacacs{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) bool) *oc.System_Aaa_ServerGroup_Server_TacacsWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_TacacsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs to the batch object.
func (n *System_Aaa_ServerGroup_Server_TacacsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Tacacs", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Tacacs_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Tacacs{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_PortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_PortPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/port to the batch object.
func (n *System_Aaa_ServerGroup_Server_Tacacs_PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath extracts the value of the leaf Port from its parent oc.System_Aaa_ServerGroup_Server_Tacacs
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_Tacacs_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Tacacs) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Port
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyHashedPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/secret-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SecretKeyPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Tacacs", gs, queryPath, true, false)
		return convertSystem_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/tacacs/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Tacacs_SourceAddressPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return convertSystem_Aaa_ServerGroup_Server_TimeoutPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_TimeoutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_TimeoutPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup", gs, queryPath, true, false)
		return convertSystem_Aaa_ServerGroup_TypePath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_SERVER_TYPE) bool) *oc.E_AaaTypes_AAA_SERVER_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_TypePath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Alarm{
			Metadata: md,
		}).SetVal(gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_AlarmPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Alarm) bool) *oc.System_AlarmWatcher {
	t.Helper()
	return watch_System_AlarmPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return convertSystem_Alarm_IdPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_IdPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return convertSystem_Alarm_ResourcePath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_ResourcePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_ResourcePath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return convertSystem_Alarm_SeverityPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_SeverityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_System_Alarm_SeverityPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Alarm", gs, queryPath, true, false)
		return convertSystem_Alarm_TextPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/alarms/alarm/state/text with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Alarm_TextPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Alarm_TextPath(t, n, timeout, predicate)
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
