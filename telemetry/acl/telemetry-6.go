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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Transport_DestinationPortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) bool) *oc.Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) bool) *oc.Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Transport_DestinationPortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Await(t testing.TB, timeout time.Duration, val oc.Acl_AclSet_AclEntry_Transport_DestinationPort_Union) *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port to the batch object.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) bool) *oc.Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Transport_DestinationPortPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/destination-port to the batch object.
func (n *Acl_AclSet_AclEntry_Transport_DestinationPortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Transport_DestinationPortPath extracts the value of the leaf DestinationPort from its parent oc.Acl_AclSet_AclEntry_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union.
func convertAcl_AclSet_AclEntry_Transport_DestinationPortPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union{
		Metadata: md,
	}
	val := parent.DestinationPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Transport_SourcePortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) bool) *oc.Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) bool) *oc.Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Transport_SourcePortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Await(t testing.TB, timeout time.Duration, val oc.Acl_AclSet_AclEntry_Transport_SourcePort_Union) *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port to the batch object.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) bool) *oc.Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Transport_SourcePortPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/source-port to the batch object.
func (n *Acl_AclSet_AclEntry_Transport_SourcePortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Transport_SourcePortPath extracts the value of the leaf SourcePort from its parent oc.Acl_AclSet_AclEntry_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union.
func convertAcl_AclSet_AclEntry_Transport_SourcePortPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union{
		Metadata: md,
	}
	val := parent.SourcePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Lookup(t testing.TB) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Get(t testing.TB) []oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Get(t testing.TB) [][]oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_PacketMatchTypes_TCP_FLAGS
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	c := &oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Transport_TcpFlagsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	w := &oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Transport_TcpFlagsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Await(t testing.TB, timeout time.Duration, val []oc.E_PacketMatchTypes_TCP_FLAGS) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags to the batch object.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	c := &oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Transport_TcpFlagsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport/state/tcp-flags to the batch object.
func (n *Acl_AclSet_AclEntry_Transport_TcpFlagsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath extracts the value of the leaf TcpFlags from its parent oc.Acl_AclSet_AclEntry_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice.
func convertAcl_AclSet_AclEntry_Transport_TcpFlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	qv := &oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice{
		Metadata: md,
	}
	val := parent.TcpFlags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/state/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a ONCE subscription.
func (n *Acl_AclSet_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet", gs, queryPath, true, false)
		return convertAcl_AclSet_DescriptionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/state/description to the batch object.
func (n *Acl_AclSet_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_DescriptionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/state/description to the batch object.
func (n *Acl_AclSet_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_DescriptionPath extracts the value of the leaf Description from its parent oc.Acl_AclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a ONCE subscription.
func (n *Acl_AclSet_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet", gs, queryPath, true, false)
		return convertAcl_AclSet_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/state/name to the batch object.
func (n *Acl_AclSet_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/state/name to the batch object.
func (n *Acl_AclSet_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_NamePath extracts the value of the leaf Name from its parent oc.Acl_AclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/state/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_TypePath) Get(t testing.TB) oc.E_Acl_ACL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a ONCE subscription.
func (n *Acl_AclSet_TypePathAny) Get(t testing.TB) []oc.E_Acl_ACL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_TYPE {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_TYPE) bool) *oc.E_Acl_ACL_TYPEWatcher {
	t.Helper()
	w := &oc.E_Acl_ACL_TYPEWatcher{}
	gs := &oc.Acl_AclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet", gs, queryPath, true, false)
		return convertAcl_AclSet_TypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Acl_ACL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_TYPE) bool) *oc.E_Acl_ACL_TYPEWatcher {
	t.Helper()
	return watch_Acl_AclSet_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Acl_ACL_TYPE) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Acl_ACL_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/state/type to the batch object.
func (n *Acl_AclSet_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_TYPE {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_TYPE) bool) *oc.E_Acl_ACL_TYPEWatcher {
	t.Helper()
	return watch_Acl_AclSet_TypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/state/type to the batch object.
func (n *Acl_AclSet_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_TypePath extracts the value of the leaf Type from its parent oc.Acl_AclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_TYPE.
func convertAcl_AclSet_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/state/counter-capability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_CounterCapabilityPath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	goStruct := &oc.Acl{}
	md, ok := oc.Lookup(t, n, "Acl", goStruct, true, false)
	if ok {
		return convertAcl_CounterCapabilityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/state/counter-capability with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_CounterCapabilityPath) Get(t testing.TB) oc.E_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/state/counter-capability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_CounterCapabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_CounterCapabilityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/state/counter-capability with a ONCE subscription.
func (n *Acl_CounterCapabilityPathAny) Get(t testing.TB) []oc.E_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_COUNTER_CAPABILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_CounterCapabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_CounterCapabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool) *oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher {
	t.Helper()
	w := &oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher{}
	gs := &oc.Acl{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl", gs, queryPath, true, false)
		return convertAcl_CounterCapabilityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_CounterCapabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool) *oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher {
	t.Helper()
	return watch_Acl_CounterCapabilityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_CounterCapabilityPath) Await(t testing.TB, timeout time.Duration, val oc.E_Acl_ACL_COUNTER_CAPABILITY) *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/state/counter-capability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/state/counter-capability to the batch object.
func (n *Acl_CounterCapabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_CounterCapabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_CounterCapabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool) *oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher {
	t.Helper()
	return watch_Acl_CounterCapabilityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/state/counter-capability to the batch object.
func (n *Acl_CounterCapabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_CounterCapabilityPath extracts the value of the leaf CounterCapability from its parent oc.Acl
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY.
func convertAcl_CounterCapabilityPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl) *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY{
		Metadata: md,
	}
	val := parent.CounterCapability
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
