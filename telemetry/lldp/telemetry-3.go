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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_FrameErrorInPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_FrameErrorInPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_FrameErrorInPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_FrameErrorInPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLldp_Interface_Counters_FrameErrorInPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a ONCE subscription.
func (n *Lldp_Interface_Counters_FrameErrorInPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameErrorInPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameErrorInPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_FrameErrorInPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameErrorInPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameErrorInPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_FrameErrorInPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in to the batch object.
func (n *Lldp_Interface_Counters_FrameErrorInPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameErrorInPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameErrorInPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Counters_FrameErrorInPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameErrorInPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameErrorInPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-in to the batch object.
func (n *Lldp_Interface_Counters_FrameErrorInPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_FrameErrorInPath extracts the value of the leaf FrameErrorIn from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_FrameErrorInPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_FrameErrorOutPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_FrameErrorOutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_FrameErrorOutPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_FrameErrorOutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLldp_Interface_Counters_FrameErrorOutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a ONCE subscription.
func (n *Lldp_Interface_Counters_FrameErrorOutPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameErrorOutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameErrorOutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_FrameErrorOutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameErrorOutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameErrorOutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_FrameErrorOutPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out to the batch object.
func (n *Lldp_Interface_Counters_FrameErrorOutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameErrorOutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameErrorOutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Counters_FrameErrorOutPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameErrorOutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameErrorOutPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-error-out to the batch object.
func (n *Lldp_Interface_Counters_FrameErrorOutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_FrameErrorOutPath extracts the value of the leaf FrameErrorOut from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_FrameErrorOutPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FrameErrorOut
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_FrameInPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_FrameInPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_FrameInPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_FrameInPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLldp_Interface_Counters_FrameInPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a ONCE subscription.
func (n *Lldp_Interface_Counters_FrameInPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameInPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameInPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_FrameInPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameInPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameInPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_FrameInPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in to the batch object.
func (n *Lldp_Interface_Counters_FrameInPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameInPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameInPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Counters_FrameInPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameInPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameInPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-in to the batch object.
func (n *Lldp_Interface_Counters_FrameInPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_FrameInPath extracts the value of the leaf FrameIn from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_FrameInPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_FrameOutPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_FrameOutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_FrameOutPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_FrameOutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLldp_Interface_Counters_FrameOutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a ONCE subscription.
func (n *Lldp_Interface_Counters_FrameOutPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameOutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameOutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_FrameOutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameOutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameOutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_FrameOutPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out to the batch object.
func (n *Lldp_Interface_Counters_FrameOutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_FrameOutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_FrameOutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Counters_FrameOutPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_FrameOutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_FrameOutPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/frame-out to the batch object.
func (n *Lldp_Interface_Counters_FrameOutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_FrameOutPath extracts the value of the leaf FrameOut from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_FrameOutPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_LastClearPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_LastClearPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_LastClearPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_LastClearPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Counters_LastClearPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a ONCE subscription.
func (n *Lldp_Interface_Counters_LastClearPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_LastClearPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_LastClearPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_LastClearPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_LastClearPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_LastClearPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_LastClearPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear to the batch object.
func (n *Lldp_Interface_Counters_LastClearPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_LastClearPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_LastClearPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Counters_LastClearPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_LastClearPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_LastClearPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/last-clear to the batch object.
func (n *Lldp_Interface_Counters_LastClearPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_LastClearPath extracts the value of the leaf LastClear from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_Counters_LastClearPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LastClear
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_TlvDiscardPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_TlvDiscardPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_TlvDiscardPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_TlvDiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLldp_Interface_Counters_TlvDiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a ONCE subscription.
func (n *Lldp_Interface_Counters_TlvDiscardPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_TlvDiscardPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_TlvDiscardPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_TlvDiscardPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_TlvDiscardPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_TlvDiscardPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_TlvDiscardPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard to the batch object.
func (n *Lldp_Interface_Counters_TlvDiscardPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_TlvDiscardPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_TlvDiscardPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Counters_TlvDiscardPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_TlvDiscardPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_TlvDiscardPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-discard to the batch object.
func (n *Lldp_Interface_Counters_TlvDiscardPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_TlvDiscardPath extracts the value of the leaf TlvDiscard from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_TlvDiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
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
