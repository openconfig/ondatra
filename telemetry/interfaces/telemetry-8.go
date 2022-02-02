package interfaces

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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_OutForwardedPktsPath extracts the value of the leaf OutForwardedPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_OutForwardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutForwardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Counters_OutOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Counters_OutPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_DhcpClientPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDhcpClient())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_DhcpClientPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_DhcpClientPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_DhcpClientPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_DhcpClientPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client to the batch object.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_DhcpClientPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dhcp-client to the batch object.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_DhcpClientPath extracts the value of the leaf DhcpClient from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_DhcpClientPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.DhcpClient
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetDupAddrDetectTransmits())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits to the batch object.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/dup-addr-detect-transmits to the batch object.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath extracts the value of the leaf DupAddrDetectTransmits from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.DupAddrDetectTransmits
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled to the batch object.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/enabled to the batch object.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_MtuPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_MtuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_MtuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_MtuPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_MtuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_MtuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_MtuPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu to the batch object.
func (n *Interface_Subinterface_Ipv6_MtuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_MtuPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/mtu to the batch object.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Mtu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Neighbor)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor) bool) *oc.Interface_Subinterface_Ipv6_NeighborWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_NeighborWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor) bool) *oc.Interface_Subinterface_Ipv6_NeighborWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor to the batch object.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor) bool) *oc.Interface_Subinterface_Ipv6_NeighborWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_NeighborPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor to the batch object.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Neighbor_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_IpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/ip to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Neighbor_IpPath extracts the value of the leaf Ip from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_IsRouterPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_IsRouterPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Neighbor_IsRouterPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Neighbor_IsRouterPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_IsRouterPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_IsRouterPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/is-router to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IsRouterPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Neighbor_IsRouterPath extracts the value of the leaf IsRouter from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Neighbor_IsRouterPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.IsRouter
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/link-layer-address to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath extracts the value of the leaf LinkLayerAddress from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LinkLayerAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath) Lookup(t testing.TB) *oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath) Get(t testing.TB) oc.E_Neighbor_NeighborState {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Neighbor_NeighborState
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePathAny) Get(t testing.TB) []oc.E_Neighbor_NeighborState {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Neighbor_NeighborState
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Neighbor_NeighborState {
	t.Helper()
	c := &oc.CollectionE_Neighbor_NeighborState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Neighbor_NeighborState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Neighbor_NeighborState) bool) *oc.E_Neighbor_NeighborStateWatcher {
	t.Helper()
	w := &oc.E_Neighbor_NeighborStateWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Neighbor_NeighborState)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Neighbor_NeighborState) bool) *oc.E_Neighbor_NeighborStateWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath) Await(t testing.TB, timeout time.Duration, val oc.E_Neighbor_NeighborState) *oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Neighbor_NeighborState) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Neighbor_NeighborState {
	t.Helper()
	c := &oc.CollectionE_Neighbor_NeighborState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Neighbor_NeighborState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Neighbor_NeighborState) bool) *oc.E_Neighbor_NeighborStateWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/neighbor-state to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_NeighborStatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Neighbor_NeighborStatePath extracts the value of the leaf NeighborState from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Neighbor_NeighborState.
func convertInterface_Subinterface_Ipv6_Neighbor_NeighborStatePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedE_Neighbor_NeighborState {
	t.Helper()
	qv := &oc.QualifiedE_Neighbor_NeighborState{
		Metadata: md,
	}
	val := parent.NeighborState
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPath) Get(t testing.TB) oc.E_IfIp_NeighborOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_NeighborOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_NeighborOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_NeighborOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_NeighborOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_NeighborOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Neighbor_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_NeighborOriginWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Neighbor_OriginPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_NeighborOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_NeighborOrigin) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_NeighborOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_NeighborOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Neighbor_OriginPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/state/origin to the batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Neighbor_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_NeighborOrigin.
func convertInterface_Subinterface_Ipv6_Neighbor_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_NeighborOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_RouterAdvertisement
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_RouterAdvertisement{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_RouterAdvertisement)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_RouterAdvertisementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) bool) *oc.Interface_Subinterface_Ipv6_RouterAdvertisementWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_RouterAdvertisementWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) bool) *oc.Interface_Subinterface_Ipv6_RouterAdvertisementWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_RouterAdvertisement{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) bool) *oc.Interface_Subinterface_Ipv6_RouterAdvertisementWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisementPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/interval to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath extracts the value of the leaf Interval from its parent oc.Interface_Subinterface_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/lifetime to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath extracts the value of the leaf Lifetime from its parent oc.Interface_Subinterface_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Lifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSuppress())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/state/suppress to the batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath extracts the value of the leaf Suppress from its parent oc.Interface_Subinterface_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Suppress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Unnumbered
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Unnumbered{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Unnumbered)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_UnnumberedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered) bool) *oc.Interface_Subinterface_Ipv6_UnnumberedWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_UnnumberedWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered) bool) *oc.Interface_Subinterface_Ipv6_UnnumberedWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_UnnumberedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Unnumbered) *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered to the batch object.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Unnumbered{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered) bool) *oc.Interface_Subinterface_Ipv6_UnnumberedWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_UnnumberedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered to the batch object.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Unnumbered_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/state/enabled to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_Subinterface_Ipv6_Unnumbered
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Unnumbered) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_LastChangePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_LastChangePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_LastChangePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_LastChangePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_LastChangePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a ONCE subscription.
func (n *Interface_Subinterface_LastChangePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_LastChangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_LastChangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return convertInterface_Subinterface_LastChangePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_LastChangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_LastChangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_LastChangePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change to the batch object.
func (n *Interface_Subinterface_LastChangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_LastChangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_LastChangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_LastChangePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/last-change to the batch object.
func (n *Interface_Subinterface_LastChangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_LastChangePath extracts the value of the leaf LastChange from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_LastChangePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastChange
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_LogicalPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_LogicalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_LogicalPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_LogicalPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_LogicalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a ONCE subscription.
func (n *Interface_Subinterface_LogicalPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_LogicalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_LogicalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return convertInterface_Subinterface_LogicalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_LogicalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_LogicalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_LogicalPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical to the batch object.
func (n *Interface_Subinterface_LogicalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_LogicalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_LogicalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_LogicalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/logical to the batch object.
func (n *Interface_Subinterface_LogicalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_LogicalPath extracts the value of the leaf Logical from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_LogicalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Logical
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_ManagementPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_ManagementPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_ManagementPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_ManagementPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_ManagementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a ONCE subscription.
func (n *Interface_Subinterface_ManagementPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_ManagementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_ManagementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return convertInterface_Subinterface_ManagementPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_ManagementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_ManagementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_ManagementPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management to the batch object.
func (n *Interface_Subinterface_ManagementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_ManagementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_ManagementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_ManagementPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/management to the batch object.
func (n *Interface_Subinterface_ManagementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_ManagementPath extracts the value of the leaf Management from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_ManagementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Management
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a ONCE subscription.
func (n *Interface_Subinterface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return convertInterface_Subinterface_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name to the batch object.
func (n *Interface_Subinterface_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/name to the batch object.
func (n *Interface_Subinterface_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_NamePath extracts the value of the leaf Name from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_OperStatusPath) Lookup(t testing.TB) *oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	goStruct := &oc.Interface_Subinterface{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_OperStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_OperStatusPath) Get(t testing.TB) oc.E_Interface_OperStatus {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_OperStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Interface_OperStatus
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_OperStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a ONCE subscription.
func (n *Interface_Subinterface_OperStatusPathAny) Get(t testing.TB) []oc.E_Interface_OperStatus {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Interface_OperStatus
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_OperStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Interface_OperStatus {
	t.Helper()
	c := &oc.CollectionE_Interface_OperStatus{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Interface_OperStatus) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_OperStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	w := &oc.E_Interface_OperStatusWatcher{}
	gs := &oc.Interface_Subinterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface", gs, queryPath, true, false)
		return convertInterface_Subinterface_OperStatusPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Interface_OperStatus)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_OperStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	return watch_Interface_Subinterface_OperStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_OperStatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_Interface_OperStatus) *oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Interface_OperStatus) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status to the batch object.
func (n *Interface_Subinterface_OperStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_OperStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Interface_OperStatus {
	t.Helper()
	c := &oc.CollectionE_Interface_OperStatus{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Interface_OperStatus) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_OperStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	return watch_Interface_Subinterface_OperStatusPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/state/oper-status to the batch object.
func (n *Interface_Subinterface_OperStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_OperStatusPath extracts the value of the leaf OperStatus from its parent oc.Interface_Subinterface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Interface_OperStatus.
func convertInterface_Subinterface_OperStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface) *oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	qv := &oc.QualifiedE_Interface_OperStatus{
		Metadata: md,
	}
	val := parent.OperStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_VlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_VlanPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_VlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription.
func (n *Interface_Subinterface_VlanPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_VlanPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_VlanPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan) bool) *oc.Interface_Subinterface_VlanWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_VlanWatcher{}
	gs := &oc.Interface_Subinterface_Vlan{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_VlanPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan) bool) *oc.Interface_Subinterface_VlanWatcher {
	t.Helper()
	return watch_Interface_Subinterface_VlanPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_VlanPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan) *oc.QualifiedInterface_Subinterface_Vlan {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan to the batch object.
func (n *Interface_Subinterface_VlanPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_VlanPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_VlanPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan) bool) *oc.Interface_Subinterface_VlanWatcher {
	t.Helper()
	return watch_Interface_Subinterface_VlanPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan to the batch object.
func (n *Interface_Subinterface_VlanPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_EgressMapping{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_EgressMapping
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_EgressMapping{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_EgressMapping
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_EgressMapping{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_EgressMapping{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_EgressMapping)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_EgressMappingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping) bool) *oc.Interface_Subinterface_Vlan_EgressMappingWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_EgressMappingWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_EgressMapping{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_EgressMapping)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping) bool) *oc.Interface_Subinterface_Vlan_EgressMappingWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMappingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_EgressMapping{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping) bool) *oc.Interface_Subinterface_Vlan_EgressMappingWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMappingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_TPID_TYPES {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_TPID_TYPES{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_EgressMapping_TpidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	w := &oc.E_VlanTypes_TPID_TYPESWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_VlanTypes_TPID_TYPES)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMapping_TpidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Await(t testing.TB, timeout time.Duration, val oc.E_VlanTypes_TPID_TYPES) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_TPID_TYPES {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_TPID_TYPES{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMapping_TpidPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/tpid to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_EgressMapping_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_EgressMapping_VlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Get(t testing.TB) oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_VlanStackAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Get(t testing.TB) []oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_VlanStackAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_VlanStackAction {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_VlanStackAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_VlanStackAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_VlanStackAction) bool) *oc.E_VlanTypes_VlanStackActionWatcher {
	t.Helper()
	w := &oc.E_VlanTypes_VlanStackActionWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_VlanTypes_VlanStackAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_VlanStackAction) bool) *oc.E_VlanTypes_VlanStackActionWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Await(t testing.TB, timeout time.Duration, val oc.E_VlanTypes_VlanStackAction) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_VlanTypes_VlanStackAction) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_VlanStackAction {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_VlanStackAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_VlanStackAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_VlanStackAction) bool) *oc.E_VlanTypes_VlanStackActionWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/state/vlan-stack-action to the batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanStackActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath extracts the value of the leaf VlanStackAction from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_VlanStackAction.
func convertInterface_Subinterface_Vlan_EgressMapping_VlanStackActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_VlanStackAction{
		Metadata: md,
	}
	val := parent.VlanStackAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_IngressMapping{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_IngressMapping
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_IngressMapping{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_IngressMapping
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_IngressMapping{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_IngressMapping{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_IngressMapping)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_IngressMappingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping) bool) *oc.Interface_Subinterface_Vlan_IngressMappingWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_IngressMappingWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_IngressMapping{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_IngressMapping)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping) bool) *oc.Interface_Subinterface_Vlan_IngressMappingWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMappingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMappingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_IngressMapping{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_IngressMapping) bool) *oc.Interface_Subinterface_Vlan_IngressMappingWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMappingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMappingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_TPID_TYPES {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_TPID_TYPES{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_IngressMapping_TpidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	w := &oc.E_VlanTypes_TPID_TYPESWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_VlanTypes_TPID_TYPES)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMapping_TpidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Await(t testing.TB, timeout time.Duration, val oc.E_VlanTypes_TPID_TYPES) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_TPID_TYPES {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_TPID_TYPES{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMapping_TpidPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/tpid to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_TpidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_IngressMapping_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface_Subinterface_Vlan_IngressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_Subinterface_Vlan_IngressMapping_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_IngressMapping_VlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMapping_VlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_IngressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_IngressMapping_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_IngressMapping", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Get(t testing.TB) oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_VlanStackAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_IngressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Get(t testing.TB) []oc.E_VlanTypes_VlanStackAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_VlanStackAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_VlanStackAction {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_VlanStackAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_VlanStackAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_VlanStackAction) bool) *oc.E_VlanTypes_VlanStackActionWatcher {
	t.Helper()
	w := &oc.E_VlanTypes_VlanStackActionWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_IngressMapping{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_IngressMapping", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_VlanTypes_VlanStackAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_VlanStackAction) bool) *oc.E_VlanTypes_VlanStackActionWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Await(t testing.TB, timeout time.Duration, val oc.E_VlanTypes_VlanStackAction) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_VlanTypes_VlanStackAction) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_VlanStackAction {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_VlanStackAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_VlanStackAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_VlanStackAction) bool) *oc.E_VlanTypes_VlanStackActionWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/ingress-mapping/state/vlan-stack-action to the batch object.
func (n *Interface_Subinterface_Vlan_IngressMapping_VlanStackActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath extracts the value of the leaf VlanStackAction from its parent oc.Interface_Subinterface_Vlan_IngressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_VlanStackAction.
func convertInterface_Subinterface_Vlan_IngressMapping_VlanStackActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_IngressMapping) *oc.QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_VlanStackAction{
		Metadata: md,
	}
	val := parent.VlanStackAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
