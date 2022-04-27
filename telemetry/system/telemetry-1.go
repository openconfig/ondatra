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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a ONCE subscription.
func (n *System_Aaa_ServerGroup_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/state/name to the batch object.
func (n *System_Aaa_ServerGroup_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertSystem_Aaa_ServerGroup_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/state/name to the batch object.
func (n *System_Aaa_ServerGroup_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_NamePath extracts the value of the leaf Name from its parent oc.System_Aaa_ServerGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_ServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup_Server {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_ServerPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup_Server {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_ServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup_Server {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup_Server
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a ONCE subscription.
func (n *System_Aaa_ServerGroup_ServerPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup_Server {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup_Server
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_ServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_ServerGroup_Server{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_ServerGroup_Server)))
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_ServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool) *oc.System_Aaa_ServerGroup_ServerWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_ServerWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_ServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool) *oc.System_Aaa_ServerGroup_ServerWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_ServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_ServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedSystem_Aaa_ServerGroup_Server {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server to the batch object.
func (n *System_Aaa_ServerGroup_ServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_ServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_ServerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool) *oc.System_Aaa_ServerGroup_ServerWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_ServerWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_ServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server) bool) *oc.System_Aaa_ServerGroup_ServerWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_ServerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server to the batch object.
func (n *System_Aaa_ServerGroup_ServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address to the batch object.
func (n *System_Aaa_ServerGroup_Server_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/address to the batch object.
func (n *System_Aaa_ServerGroup_Server_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_AddressPath extracts the value of the leaf Address from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_ConnectionAbortsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_ConnectionAbortsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionAbortsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_ConnectionAbortsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionAbortsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_ConnectionAbortsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-aborts to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionAbortsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_ConnectionAbortsPath extracts the value of the leaf ConnectionAborts from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_ConnectionAbortsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConnectionAborts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_ConnectionClosesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_ConnectionClosesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionClosesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_ConnectionClosesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionClosesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionClosesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_ConnectionClosesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionClosesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-closes to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionClosesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_ConnectionClosesPath extracts the value of the leaf ConnectionCloses from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_ConnectionClosesPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConnectionCloses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_ConnectionFailuresPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_ConnectionFailuresPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionFailuresPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_ConnectionFailuresPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionFailuresPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_ConnectionFailuresPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-failures to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionFailuresPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_ConnectionFailuresPath extracts the value of the leaf ConnectionFailures from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_ConnectionFailuresPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConnectionFailures
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_ConnectionOpensPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_ConnectionOpensPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionOpensPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_ConnectionOpensPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionOpensPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionOpensPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_ConnectionOpensPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionOpensPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-opens to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionOpensPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_ConnectionOpensPath extracts the value of the leaf ConnectionOpens from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_ConnectionOpensPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConnectionOpens
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/connection-timeouts to the batch object.
func (n *System_Aaa_ServerGroup_Server_ConnectionTimeoutsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_ConnectionTimeoutsPath extracts the value of the leaf ConnectionTimeouts from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_ConnectionTimeoutsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConnectionTimeouts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_ErrorsReceivedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_ErrorsReceivedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ErrorsReceivedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_ErrorsReceivedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ErrorsReceivedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received to the batch object.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_ErrorsReceivedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/errors-received to the batch object.
func (n *System_Aaa_ServerGroup_Server_ErrorsReceivedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_ErrorsReceivedPath extracts the value of the leaf ErrorsReceived from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_ErrorsReceivedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ErrorsReceived
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_MessagesReceivedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_MessagesReceivedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_MessagesReceivedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_MessagesReceivedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_MessagesReceivedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received to the batch object.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_MessagesReceivedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_MessagesReceivedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_MessagesReceivedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-received to the batch object.
func (n *System_Aaa_ServerGroup_Server_MessagesReceivedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_MessagesReceivedPath extracts the value of the leaf MessagesReceived from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_MessagesReceivedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MessagesReceived
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_MessagesSentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_MessagesSentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_MessagesSentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_MessagesSentPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_MessagesSentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent to the batch object.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_MessagesSentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_MessagesSentPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_MessagesSentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/messages-sent to the batch object.
func (n *System_Aaa_ServerGroup_Server_MessagesSentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_MessagesSentPath extracts the value of the leaf MessagesSent from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_MessagesSentPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MessagesSent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name to the batch object.
func (n *System_Aaa_ServerGroup_Server_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertSystem_Aaa_ServerGroup_Server_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/state/name to the batch object.
func (n *System_Aaa_ServerGroup_Server_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_NamePath extracts the value of the leaf Name from its parent oc.System_Aaa_ServerGroup_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup_Server_Radius
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server_Radius{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_ServerGroup_Server_Radius)))
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_RadiusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool) *oc.System_Aaa_ServerGroup_Server_RadiusWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_Server_RadiusWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool) *oc.System_Aaa_ServerGroup_Server_RadiusWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_RadiusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius to the batch object.
func (n *System_Aaa_ServerGroup_Server_RadiusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server_Radius{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_RadiusPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool) *oc.System_Aaa_ServerGroup_Server_RadiusWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_Server_RadiusWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius) bool) *oc.System_Aaa_ServerGroup_Server_RadiusWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_RadiusPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius to the batch object.
func (n *System_Aaa_ServerGroup_Server_RadiusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAcctPort())
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_AcctPortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/acct-port to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AcctPortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath extracts the value of the leaf AcctPort from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_Radius_AcctPortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AcctPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAuthPort())
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_AuthPortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/auth-port to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_AuthPortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath extracts the value of the leaf AuthPort from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Aaa_ServerGroup_Server_Radius_AuthPortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AuthPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup_Server_Radius_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_ServerGroup_Server_Radius_Counters)))
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool) *oc.System_Aaa_ServerGroup_Server_Radius_CountersWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_Server_Radius_CountersWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool) *oc.System_Aaa_ServerGroup_Server_Radius_CountersWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_ServerGroup_Server_Radius_Counters) *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool) *oc.System_Aaa_ServerGroup_Server_Radius_CountersWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroup_Server_Radius_CountersWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) bool) *oc.System_Aaa_ServerGroup_Server_Radius_CountersWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-accepts to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath extracts the value of the leaf AccessAccepts from its parent oc.System_Aaa_ServerGroup_Server_Radius_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessAcceptsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AccessAccepts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/access-rejects to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath extracts the value of the leaf AccessRejects from its parent oc.System_Aaa_ServerGroup_Server_Radius_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_Radius_Counters_AccessRejectsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AccessRejects
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/retried-access-requests to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath extracts the value of the leaf RetriedAccessRequests from its parent oc.System_Aaa_ServerGroup_Server_Radius_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_Radius_Counters_RetriedAccessRequestsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.RetriedAccessRequests
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius_Counters", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/counters/timeout-access-requests to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath extracts the value of the leaf TimeoutAccessRequests from its parent oc.System_Aaa_ServerGroup_Server_Radius_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_ServerGroup_Server_Radius_Counters_TimeoutAccessRequestsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TimeoutAccessRequests
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/retransmit-attempts to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath extracts the value of the leaf RetransmitAttempts from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Aaa_ServerGroup_Server_Radius_RetransmitAttemptsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.RetransmitAttempts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key-hashed to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath extracts the value of the leaf SecretKeyHashed from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/secret-key to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SecretKeyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath extracts the value of the leaf SecretKey from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Radius_SecretKeyPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup_Server_Radius", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup_Server_Radius{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a ONCE subscription.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup_Server_Radius{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup_Server_Radius{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup_Server_Radius", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group/servers/server/radius/state/source-address to the batch object.
func (n *System_Aaa_ServerGroup_Server_Radius_SourceAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.System_Aaa_ServerGroup_Server_Radius
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_ServerGroup_Server_Radius_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_ServerGroup_Server_Radius) *oc.QualifiedString {
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
