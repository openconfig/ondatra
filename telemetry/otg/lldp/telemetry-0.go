package lldp

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry/otg"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterfacePath) Lookup(t testing.TB) *oc.QualifiedLldpInterface {
	t.Helper()
	goStruct := &oc.LldpInterface{}
	md, ok := oc.Lookup(t, n, "LldpInterface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldpInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterfacePath) Get(t testing.TB) *oc.LldpInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedLldpInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldpInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldpInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface with a ONCE subscription.
func (n *LldpInterfacePathAny) Get(t testing.TB) []*oc.LldpInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LldpInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface {
	t.Helper()
	c := &oc.CollectionLldpInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldpInterface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LldpInterface)))
		return false
	})
	return c
}

func watch_LldpInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface) bool) *oc.LldpInterfaceWatcher {
	t.Helper()
	w := &oc.LldpInterfaceWatcher{}
	gs := &oc.LldpInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldpInterface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface) bool) *oc.LldpInterfaceWatcher {
	t.Helper()
	return watch_LldpInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.LldpInterface) *oc.QualifiedLldpInterface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldpInterface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface to the batch object.
func (n *LldpInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface {
	t.Helper()
	c := &oc.CollectionLldpInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface) bool) *oc.LldpInterfaceWatcher {
	t.Helper()
	w := &oc.LldpInterfaceWatcher{}
	structs := map[string]*oc.LldpInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldpInterface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface) bool) *oc.LldpInterfaceWatcher {
	t.Helper()
	return watch_LldpInterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface to the batch object.
func (n *LldpInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_CountersPath) Lookup(t testing.TB) *oc.QualifiedLldpInterface_Counters {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldpInterface_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_CountersPath) Get(t testing.TB) *oc.LldpInterface_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedLldpInterface_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldpInterface_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldpInterface_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a ONCE subscription.
func (n *LldpInterface_CountersPathAny) Get(t testing.TB) []*oc.LldpInterface_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LldpInterface_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_Counters {
	t.Helper()
	c := &oc.CollectionLldpInterface_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldpInterface_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LldpInterface_Counters)))
		return false
	})
	return c
}

func watch_LldpInterface_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_Counters) bool) *oc.LldpInterface_CountersWatcher {
	t.Helper()
	w := &oc.LldpInterface_CountersWatcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldpInterface_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_Counters) bool) *oc.LldpInterface_CountersWatcher {
	t.Helper()
	return watch_LldpInterface_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.LldpInterface_Counters) *oc.QualifiedLldpInterface_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldpInterface_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters to the batch object.
func (n *LldpInterface_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_Counters {
	t.Helper()
	c := &oc.CollectionLldpInterface_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_Counters) bool) *oc.LldpInterface_CountersWatcher {
	t.Helper()
	w := &oc.LldpInterface_CountersWatcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldpInterface_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_Counters) bool) *oc.LldpInterface_CountersWatcher {
	t.Helper()
	return watch_LldpInterface_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters to the batch object.
func (n *LldpInterface_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_Counters_FrameDiscardPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, true, false)
	if ok {
		return convertLldpInterface_Counters_FrameDiscardPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_Counters_FrameDiscardPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_Counters_FrameDiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_Counters_FrameDiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a ONCE subscription.
func (n *LldpInterface_Counters_FrameDiscardPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameDiscardPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameDiscardPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_Counters_FrameDiscardPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameDiscardPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameDiscardPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_Counters_FrameDiscardPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard to the batch object.
func (n *LldpInterface_Counters_FrameDiscardPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameDiscardPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameDiscardPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_Counters_FrameDiscardPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameDiscardPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameDiscardPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-discard to the batch object.
func (n *LldpInterface_Counters_FrameDiscardPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_Counters_FrameDiscardPath extracts the value of the leaf FrameDiscard from its parent oc.LldpInterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldpInterface_Counters_FrameDiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_Counters_FrameErrorInPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, true, false)
	if ok {
		return convertLldpInterface_Counters_FrameErrorInPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_Counters_FrameErrorInPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_Counters_FrameErrorInPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_Counters_FrameErrorInPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a ONCE subscription.
func (n *LldpInterface_Counters_FrameErrorInPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameErrorInPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameErrorInPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_Counters_FrameErrorInPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameErrorInPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameErrorInPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_Counters_FrameErrorInPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in to the batch object.
func (n *LldpInterface_Counters_FrameErrorInPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameErrorInPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameErrorInPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_Counters_FrameErrorInPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameErrorInPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameErrorInPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-error-in to the batch object.
func (n *LldpInterface_Counters_FrameErrorInPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_Counters_FrameErrorInPath extracts the value of the leaf FrameErrorIn from its parent oc.LldpInterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldpInterface_Counters_FrameErrorInPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FrameErrorIn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_Counters_FrameInPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, true, false)
	if ok {
		return convertLldpInterface_Counters_FrameInPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_Counters_FrameInPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_Counters_FrameInPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_Counters_FrameInPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a ONCE subscription.
func (n *LldpInterface_Counters_FrameInPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameInPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameInPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_Counters_FrameInPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameInPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameInPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_Counters_FrameInPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in to the batch object.
func (n *LldpInterface_Counters_FrameInPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameInPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameInPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_Counters_FrameInPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameInPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameInPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-in to the batch object.
func (n *LldpInterface_Counters_FrameInPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_Counters_FrameInPath extracts the value of the leaf FrameIn from its parent oc.LldpInterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldpInterface_Counters_FrameInPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FrameIn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_Counters_FrameOutPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, true, false)
	if ok {
		return convertLldpInterface_Counters_FrameOutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_Counters_FrameOutPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_Counters_FrameOutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_Counters_FrameOutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a ONCE subscription.
func (n *LldpInterface_Counters_FrameOutPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameOutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameOutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_Counters_FrameOutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameOutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameOutPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_Counters_FrameOutPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out to the batch object.
func (n *LldpInterface_Counters_FrameOutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_FrameOutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_FrameOutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_Counters_FrameOutPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_FrameOutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_FrameOutPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/frame-out to the batch object.
func (n *LldpInterface_Counters_FrameOutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_Counters_FrameOutPath extracts the value of the leaf FrameOut from its parent oc.LldpInterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldpInterface_Counters_FrameOutPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FrameOut
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_Counters_TlvDiscardPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, true, false)
	if ok {
		return convertLldpInterface_Counters_TlvDiscardPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_Counters_TlvDiscardPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_Counters_TlvDiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_Counters_TlvDiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a ONCE subscription.
func (n *LldpInterface_Counters_TlvDiscardPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_TlvDiscardPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_TlvDiscardPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_Counters_TlvDiscardPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_TlvDiscardPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_TlvDiscardPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_Counters_TlvDiscardPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard to the batch object.
func (n *LldpInterface_Counters_TlvDiscardPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_TlvDiscardPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_TlvDiscardPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_Counters_TlvDiscardPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_TlvDiscardPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_TlvDiscardPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-discard to the batch object.
func (n *LldpInterface_Counters_TlvDiscardPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_Counters_TlvDiscardPath extracts the value of the leaf TlvDiscard from its parent oc.LldpInterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldpInterface_Counters_TlvDiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TlvDiscard
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_Counters_TlvUnknownPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.LldpInterface_Counters{}
	md, ok := oc.Lookup(t, n, "LldpInterface_Counters", goStruct, true, false)
	if ok {
		return convertLldpInterface_Counters_TlvUnknownPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_Counters_TlvUnknownPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_Counters_TlvUnknownPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_Counters_TlvUnknownPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a ONCE subscription.
func (n *LldpInterface_Counters_TlvUnknownPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_TlvUnknownPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_TlvUnknownPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_Counters_TlvUnknownPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_TlvUnknownPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_TlvUnknownPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_Counters_TlvUnknownPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown to the batch object.
func (n *LldpInterface_Counters_TlvUnknownPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_Counters_TlvUnknownPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_Counters_TlvUnknownPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.LldpInterface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_Counters_TlvUnknownPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_Counters_TlvUnknownPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_LldpInterface_Counters_TlvUnknownPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/counters/tlv-unknown to the batch object.
func (n *LldpInterface_Counters_TlvUnknownPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_Counters_TlvUnknownPath extracts the value of the leaf TlvUnknown from its parent oc.LldpInterface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldpInterface_Counters_TlvUnknownPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TlvUnknown
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabasePath) Lookup(t testing.TB) *oc.QualifiedLldpInterface_LldpNeighborDatabase {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldpInterface_LldpNeighborDatabase{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabasePath) Get(t testing.TB) *oc.LldpInterface_LldpNeighborDatabase {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabasePathAny) Lookup(t testing.TB) []*oc.QualifiedLldpInterface_LldpNeighborDatabase {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldpInterface_LldpNeighborDatabase
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabasePathAny) Get(t testing.TB) []*oc.LldpInterface_LldpNeighborDatabase {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LldpInterface_LldpNeighborDatabase
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabasePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldpInterface_LldpNeighborDatabase{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LldpInterface_LldpNeighborDatabase)))
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabasePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool) *oc.LldpInterface_LldpNeighborDatabaseWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabaseWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabasePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool) *oc.LldpInterface_LldpNeighborDatabaseWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabasePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabasePath) Await(t testing.TB, timeout time.Duration, val *oc.LldpInterface_LldpNeighborDatabase) *oc.QualifiedLldpInterface_LldpNeighborDatabase {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database to the batch object.
func (n *LldpInterface_LldpNeighborDatabasePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabasePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabasePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool) *oc.LldpInterface_LldpNeighborDatabaseWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabaseWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabasePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase) bool) *oc.LldpInterface_LldpNeighborDatabaseWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabasePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database to the batch object.
func (n *LldpInterface_LldpNeighborDatabasePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPath) Lookup(t testing.TB) *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPath) Get(t testing.TB) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny) Get(t testing.TB) []*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor)))
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighborPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetAge())
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/age to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath extracts the value of the leaf Age from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_AgePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Age
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
