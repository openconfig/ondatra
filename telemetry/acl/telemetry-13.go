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

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, true, false)
		return convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath extracts the value of the leaf MatchedOctets from its parent oc.Acl_Interface_EgressAclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, true, false)
		return convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath extracts the value of the leaf MatchedPackets from its parent oc.Acl_Interface_EgressAclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedPackets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, true, false)
		return convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath extracts the value of the leaf SequenceId from its parent oc.Acl_Interface_EgressAclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SequenceId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_SetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_SetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_SetNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet", gs, queryPath, true, false)
		return convertAcl_Interface_EgressAclSet_SetNamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_SetNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name to the batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_SetNamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name to the batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_SetNamePath extracts the value of the leaf SetName from its parent oc.Acl_Interface_EgressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_EgressAclSet_SetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
