package acl

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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Lookup(t testing.TB) *oc.QualifiedUint8Slice {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Get(t testing.TB) []uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Get(t testing.TB) [][]uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8Slice {
	t.Helper()
	c := &oc.CollectionUint8Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_DscpSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8Slice) bool) *oc.Uint8SliceWatcher {
	t.Helper()
	w := &oc.Uint8SliceWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8Slice) bool) *oc.Uint8SliceWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_DscpSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Await(t testing.TB, timeout time.Duration, val []uint8) *oc.QualifiedUint8Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8Slice {
	t.Helper()
	c := &oc.CollectionUint8Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8Slice) bool) *oc.Uint8SliceWatcher {
	t.Helper()
	w := &oc.Uint8SliceWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv4{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8Slice) bool) *oc.Uint8SliceWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp-set to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath extracts the value of the leaf DscpSet from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8Slice.
func convertAcl_AclSet_AclEntry_Ipv4_DscpSetPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedUint8Slice {
	t.Helper()
	qv := &oc.QualifiedUint8Slice{
		Metadata: md,
	}
	val := parent.DscpSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_HopLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_HopLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv4{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/hop-limit to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_HopLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath extracts the value of the leaf HopLimit from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Ipv4_HopLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.HopLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_ProtocolPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool) *oc.Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool) *oc.Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_ProtocolPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Await(t testing.TB, timeout time.Duration, val oc.Acl_AclSet_AclEntry_Ipv4_Protocol_Union) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool) *oc.Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv4{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) bool) *oc.Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/protocol to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_ProtocolPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath extracts the value of the leaf Protocol from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union.
func convertAcl_AclSet_AclEntry_Ipv4_ProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union{
		Metadata: md,
	}
	val := parent.Protocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_SourceAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv4{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/source-address to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_SourceAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv4_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Ipv6{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_Ipv6)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv6Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool) *oc.Acl_AclSet_AclEntry_Ipv6Watcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Ipv6Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", gs, queryPath, false, false)
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Ipv6)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool) *oc.Acl_AclSet_AclEntry_Ipv6Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv6Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv6Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Ipv6{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv6PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool) *oc.Acl_AclSet_AclEntry_Ipv6Watcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Ipv6Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv6{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedAcl_AclSet_AclEntry_Ipv6{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Ipv6)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv6) bool) *oc.Acl_AclSet_AclEntry_Ipv6Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv6PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6 to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv6PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv6{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-address to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv6_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv6", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Ipv6{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv6", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv6/state/destination-flow-label to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath extracts the value of the leaf DestinationFlowLabel from its parent oc.Acl_AclSet_AclEntry_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_Ipv6_DestinationFlowLabelPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.DestinationFlowLabel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
