package lldp

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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_InterfacePath) Lookup(t testing.TB) *oc.QualifiedLldp_Interface {
	t.Helper()
	goStruct := &oc.Lldp_Interface{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_InterfacePath) Get(t testing.TB) *oc.Lldp_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription.
func (n *Lldp_InterfacePathAny) Get(t testing.TB) []*oc.Lldp_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface {
	t.Helper()
	c := &oc.CollectionLldp_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp_Interface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp_Interface)))
		return false
	})
	return c
}

func watch_Lldp_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface) bool) *oc.Lldp_InterfaceWatcher {
	t.Helper()
	w := &oc.Lldp_InterfaceWatcher{}
	gs := &oc.Lldp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface", gs, queryPath, false, false)
		return (&oc.QualifiedLldp_Interface{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface) bool) *oc.Lldp_InterfaceWatcher {
	t.Helper()
	return watch_Lldp_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp_Interface) *oc.QualifiedLldp_Interface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp_Interface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface to the batch object.
func (n *Lldp_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface {
	t.Helper()
	c := &oc.CollectionLldp_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface) bool) *oc.Lldp_InterfaceWatcher {
	t.Helper()
	return watch_Lldp_InterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface to the batch object.
func (n *Lldp_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_CountersPath) Lookup(t testing.TB) *oc.QualifiedLldp_Interface_Counters {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp_Interface_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_CountersPath) Get(t testing.TB) *oc.Lldp_Interface_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Interface_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Interface_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Interface_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a ONCE subscription.
func (n *Lldp_Interface_CountersPathAny) Get(t testing.TB) []*oc.Lldp_Interface_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Interface_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Counters {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp_Interface_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp_Interface_Counters)))
		return false
	})
	return c
}

func watch_Lldp_Interface_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Counters) bool) *oc.Lldp_Interface_CountersWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_CountersWatcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedLldp_Interface_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Counters) bool) *oc.Lldp_Interface_CountersWatcher {
	t.Helper()
	return watch_Lldp_Interface_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp_Interface_Counters) *oc.QualifiedLldp_Interface_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp_Interface_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters to the batch object.
func (n *Lldp_Interface_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Counters {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Counters) bool) *oc.Lldp_Interface_CountersWatcher {
	t.Helper()
	return watch_Lldp_Interface_CountersPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters to the batch object.
func (n *Lldp_Interface_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_FrameDiscardPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_FrameDiscardPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_FrameDiscardPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_FrameDiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Counters_FrameDiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a ONCE subscription.
func (n *Lldp_Interface_Counters_FrameDiscardPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameDiscardPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameDiscardPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return convertLldp_Interface_Counters_FrameDiscardPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameDiscardPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameDiscardPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_FrameDiscardPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard to the batch object.
func (n *Lldp_Interface_Counters_FrameDiscardPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameDiscardPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameDiscardPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameDiscardPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-discard to the batch object.
func (n *Lldp_Interface_Counters_FrameDiscardPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_FrameDiscardPath extracts the value of the leaf FrameDiscard from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_FrameDiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FrameDiscard
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
