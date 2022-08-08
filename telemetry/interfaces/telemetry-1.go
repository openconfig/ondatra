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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_CountersPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ethernet_Counters {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Ethernet_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_CountersPath) Get(t testing.TB) *oc.Interface_Ethernet_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ethernet_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ethernet_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Ethernet_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a ONCE subscription.
func (n *Interface_Ethernet_CountersPathAny) Get(t testing.TB) []*oc.Interface_Ethernet_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Ethernet_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ethernet_Counters {
	t.Helper()
	c := &oc.CollectionInterface_Ethernet_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ethernet_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Ethernet_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Ethernet_Counters)))
		return false
	})
	return c
}

func watch_Interface_Ethernet_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters) bool) *oc.Interface_Ethernet_CountersWatcher {
	t.Helper()
	w := &oc.Interface_Ethernet_CountersWatcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Ethernet_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ethernet_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters) bool) *oc.Interface_Ethernet_CountersWatcher {
	t.Helper()
	return watch_Interface_Ethernet_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Ethernet_Counters) *oc.QualifiedInterface_Ethernet_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Ethernet_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters to the batch object.
func (n *Interface_Ethernet_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ethernet_Counters {
	t.Helper()
	c := &oc.CollectionInterface_Ethernet_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ethernet_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters) bool) *oc.Interface_Ethernet_CountersWatcher {
	t.Helper()
	w := &oc.Interface_Ethernet_CountersWatcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Ethernet_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ethernet_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters) bool) *oc.Interface_Ethernet_CountersWatcher {
	t.Helper()
	return watch_Interface_Ethernet_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters to the batch object.
func (n *Interface_Ethernet_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InBlockErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InBlockErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InBlockErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InBlockErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InBlockErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InBlockErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InBlockErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InBlockErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InBlockErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InBlockErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InBlockErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InBlockErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors to the batch object.
func (n *Interface_Ethernet_Counters_InBlockErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InBlockErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InBlockErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InBlockErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InBlockErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InBlockErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-block-errors to the batch object.
func (n *Interface_Ethernet_Counters_InBlockErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InBlockErrorsPath extracts the value of the leaf InBlockErrors from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InBlockErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InBlockErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InCarrierErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InCarrierErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InCarrierErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InCarrierErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InCarrierErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors to the batch object.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InCarrierErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InCarrierErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InCarrierErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-carrier-errors to the batch object.
func (n *Interface_Ethernet_Counters_InCarrierErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InCarrierErrorsPath extracts the value of the leaf InCarrierErrors from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InCarrierErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InCarrierErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InCrcErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InCrcErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InCrcErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InCrcErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InCrcErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InCrcErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InCrcErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InCrcErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InCrcErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InCrcErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InCrcErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InCrcErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors to the batch object.
func (n *Interface_Ethernet_Counters_InCrcErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InCrcErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InCrcErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InCrcErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InCrcErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InCrcErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-crc-errors to the batch object.
func (n *Interface_Ethernet_Counters_InCrcErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InCrcErrorsPath extracts the value of the leaf InCrcErrors from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InCrcErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InCrcErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistributionPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ethernet_Counters_InDistribution {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Ethernet_Counters_InDistribution{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistributionPath) Get(t testing.TB) *oc.Interface_Ethernet_Counters_InDistribution {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistributionPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ethernet_Counters_InDistribution {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ethernet_Counters_InDistribution
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Ethernet_Counters_InDistribution{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistributionPathAny) Get(t testing.TB) []*oc.Interface_Ethernet_Counters_InDistribution {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Ethernet_Counters_InDistribution
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistributionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ethernet_Counters_InDistribution {
	t.Helper()
	c := &oc.CollectionInterface_Ethernet_Counters_InDistribution{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Ethernet_Counters_InDistribution{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Ethernet_Counters_InDistribution)))
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistributionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool) *oc.Interface_Ethernet_Counters_InDistributionWatcher {
	t.Helper()
	w := &oc.Interface_Ethernet_Counters_InDistributionWatcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Ethernet_Counters_InDistribution{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ethernet_Counters_InDistribution)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistributionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool) *oc.Interface_Ethernet_Counters_InDistributionWatcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistributionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistributionPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedInterface_Ethernet_Counters_InDistribution {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution to the batch object.
func (n *Interface_Ethernet_Counters_InDistributionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistributionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ethernet_Counters_InDistribution {
	t.Helper()
	c := &oc.CollectionInterface_Ethernet_Counters_InDistribution{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistributionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool) *oc.Interface_Ethernet_Counters_InDistributionWatcher {
	t.Helper()
	w := &oc.Interface_Ethernet_Counters_InDistributionWatcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Ethernet_Counters_InDistribution{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ethernet_Counters_InDistribution)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistributionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ethernet_Counters_InDistribution) bool) *oc.Interface_Ethernet_Counters_InDistributionWatcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistributionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution to the batch object.
func (n *Interface_Ethernet_Counters_InDistributionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-1024-1518-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath extracts the value of the leaf InFrames_1024_1518Octets from its parent oc.Interface_Ethernet_Counters_InDistribution
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InDistribution_InFrames_1024_1518OctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames_1024_1518Octets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-128-255-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath extracts the value of the leaf InFrames_128_255Octets from its parent oc.Interface_Ethernet_Counters_InDistribution
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InDistribution_InFrames_128_255OctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames_128_255Octets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-256-511-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath extracts the value of the leaf InFrames_256_511Octets from its parent oc.Interface_Ethernet_Counters_InDistribution
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InDistribution_InFrames_256_511OctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames_256_511Octets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-512-1023-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath extracts the value of the leaf InFrames_512_1023Octets from its parent oc.Interface_Ethernet_Counters_InDistribution
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InDistribution_InFrames_512_1023OctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames_512_1023Octets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-64-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_64OctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath extracts the value of the leaf InFrames_64Octets from its parent oc.Interface_Ethernet_Counters_InDistribution
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InDistribution_InFrames_64OctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames_64Octets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters_InDistribution", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters_InDistribution{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters_InDistribution{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters_InDistribution{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters_InDistribution", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-distribution/in-frames-65-127-octets to the batch object.
func (n *Interface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath extracts the value of the leaf InFrames_65_127Octets from its parent oc.Interface_Ethernet_Counters_InDistribution
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InDistribution_InFrames_65_127OctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters_InDistribution) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames_65_127Octets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InFragmentFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InFragmentFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InFragmentFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InFragmentFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InFragmentFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InFragmentFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InFragmentFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InFragmentFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InFragmentFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InFragmentFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InFragmentFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InFragmentFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames to the batch object.
func (n *Interface_Ethernet_Counters_InFragmentFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InFragmentFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InFragmentFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InFragmentFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InFragmentFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InFragmentFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-fragment-frames to the batch object.
func (n *Interface_Ethernet_Counters_InFragmentFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InFragmentFramesPath extracts the value of the leaf InFragmentFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InFragmentFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFragmentFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InInterruptedTxPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InInterruptedTxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InInterruptedTxPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InInterruptedTxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InInterruptedTxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InInterruptedTxPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InInterruptedTxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InInterruptedTxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InInterruptedTxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InInterruptedTxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InInterruptedTxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InInterruptedTxPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx to the batch object.
func (n *Interface_Ethernet_Counters_InInterruptedTxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InInterruptedTxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InInterruptedTxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InInterruptedTxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InInterruptedTxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InInterruptedTxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-interrupted-tx to the batch object.
func (n *Interface_Ethernet_Counters_InInterruptedTxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InInterruptedTxPath extracts the value of the leaf InInterruptedTx from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InInterruptedTxPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InInterruptedTx
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InJabberFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InJabberFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InJabberFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InJabberFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InJabberFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InJabberFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InJabberFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InJabberFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InJabberFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InJabberFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InJabberFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InJabberFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames to the batch object.
func (n *Interface_Ethernet_Counters_InJabberFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InJabberFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InJabberFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InJabberFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InJabberFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InJabberFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-jabber-frames to the batch object.
func (n *Interface_Ethernet_Counters_InJabberFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InJabberFramesPath extracts the value of the leaf InJabberFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InJabberFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InJabberFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InLateCollisionPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InLateCollisionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InLateCollisionPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InLateCollisionPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InLateCollisionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InLateCollisionPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InLateCollisionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InLateCollisionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InLateCollisionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InLateCollisionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InLateCollisionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InLateCollisionPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision to the batch object.
func (n *Interface_Ethernet_Counters_InLateCollisionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InLateCollisionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InLateCollisionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InLateCollisionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InLateCollisionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InLateCollisionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-late-collision to the batch object.
func (n *Interface_Ethernet_Counters_InLateCollisionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InLateCollisionPath extracts the value of the leaf InLateCollision from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InLateCollisionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InLateCollision
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InMacControlFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InMacControlFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InMacControlFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InMacControlFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InMacControlFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InMacControlFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMacControlFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMacControlFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InMacControlFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMacControlFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMacControlFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InMacControlFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames to the batch object.
func (n *Interface_Ethernet_Counters_InMacControlFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMacControlFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMacControlFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InMacControlFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMacControlFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMacControlFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-control-frames to the batch object.
func (n *Interface_Ethernet_Counters_InMacControlFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InMacControlFramesPath extracts the value of the leaf InMacControlFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InMacControlFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InMacControlFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InMacErrorsRxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InMacErrorsRxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMacErrorsRxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InMacErrorsRxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMacErrorsRxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx to the batch object.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMacErrorsRxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InMacErrorsRxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMacErrorsRxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-errors-rx to the batch object.
func (n *Interface_Ethernet_Counters_InMacErrorsRxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InMacErrorsRxPath extracts the value of the leaf InMacErrorsRx from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InMacErrorsRxPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InMacErrorsRx
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InMacPauseFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InMacPauseFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMacPauseFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InMacPauseFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMacPauseFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames to the batch object.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMacPauseFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InMacPauseFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMacPauseFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-mac-pause-frames to the batch object.
func (n *Interface_Ethernet_Counters_InMacPauseFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InMacPauseFramesPath extracts the value of the leaf InMacPauseFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InMacPauseFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InMacPauseFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InMaxsizeExceededPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InMaxsizeExceededPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMaxsizeExceededPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InMaxsizeExceededPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMaxsizeExceededPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded to the batch object.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InMaxsizeExceededPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InMaxsizeExceededPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InMaxsizeExceededPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-maxsize-exceeded to the batch object.
func (n *Interface_Ethernet_Counters_InMaxsizeExceededPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InMaxsizeExceededPath extracts the value of the leaf InMaxsizeExceeded from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InMaxsizeExceededPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InMaxsizeExceeded
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InOversizeFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InOversizeFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InOversizeFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InOversizeFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InOversizeFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InOversizeFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InOversizeFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InOversizeFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InOversizeFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InOversizeFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InOversizeFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InOversizeFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames to the batch object.
func (n *Interface_Ethernet_Counters_InOversizeFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InOversizeFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InOversizeFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InOversizeFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InOversizeFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InOversizeFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-oversize-frames to the batch object.
func (n *Interface_Ethernet_Counters_InOversizeFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InOversizeFramesPath extracts the value of the leaf InOversizeFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InOversizeFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOversizeFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InSingleCollisionPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InSingleCollisionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InSingleCollisionPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InSingleCollisionPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InSingleCollisionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InSingleCollisionPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InSingleCollisionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InSingleCollisionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InSingleCollisionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InSingleCollisionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InSingleCollisionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InSingleCollisionPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision to the batch object.
func (n *Interface_Ethernet_Counters_InSingleCollisionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InSingleCollisionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InSingleCollisionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InSingleCollisionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InSingleCollisionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InSingleCollisionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-single-collision to the batch object.
func (n *Interface_Ethernet_Counters_InSingleCollisionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InSingleCollisionPath extracts the value of the leaf InSingleCollision from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InSingleCollisionPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InSingleCollision
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InSymbolErrorPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InSymbolErrorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InSymbolErrorPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InSymbolErrorPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InSymbolErrorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InSymbolErrorPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InSymbolErrorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InSymbolErrorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InSymbolErrorPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InSymbolErrorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InSymbolErrorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InSymbolErrorPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error to the batch object.
func (n *Interface_Ethernet_Counters_InSymbolErrorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InSymbolErrorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InSymbolErrorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InSymbolErrorPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InSymbolErrorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InSymbolErrorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-symbol-error to the batch object.
func (n *Interface_Ethernet_Counters_InSymbolErrorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InSymbolErrorPath extracts the value of the leaf InSymbolError from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InSymbolErrorPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InSymbolError
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_InUndersizeFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_InUndersizeFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InUndersizeFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_InUndersizeFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InUndersizeFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames to the batch object.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_InUndersizeFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_InUndersizeFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_InUndersizeFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-undersize-frames to the batch object.
func (n *Interface_Ethernet_Counters_InUndersizeFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_InUndersizeFramesPath extracts the value of the leaf InUndersizeFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_InUndersizeFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InUndersizeFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_In_8021QFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_In_8021QFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_In_8021QFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_In_8021QFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_In_8021QFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_In_8021QFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_In_8021QFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_In_8021QFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_In_8021QFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_In_8021QFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_In_8021QFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_In_8021QFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames to the batch object.
func (n *Interface_Ethernet_Counters_In_8021QFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_In_8021QFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_In_8021QFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_In_8021QFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_In_8021QFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_In_8021QFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/in-8021q-frames to the batch object.
func (n *Interface_Ethernet_Counters_In_8021QFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_In_8021QFramesPath extracts the value of the leaf In_8021QFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_In_8021QFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.In_8021QFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_OutMacControlFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_OutMacControlFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_OutMacControlFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_OutMacControlFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_OutMacControlFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames to the batch object.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_OutMacControlFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_OutMacControlFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_OutMacControlFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-control-frames to the batch object.
func (n *Interface_Ethernet_Counters_OutMacControlFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_OutMacControlFramesPath extracts the value of the leaf OutMacControlFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_OutMacControlFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMacControlFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_OutMacErrorsTxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_OutMacErrorsTxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a ONCE subscription.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_OutMacErrorsTxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_OutMacErrorsTxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_OutMacErrorsTxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx to the batch object.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_OutMacErrorsTxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_OutMacErrorsTxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_OutMacErrorsTxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-errors-tx to the batch object.
func (n *Interface_Ethernet_Counters_OutMacErrorsTxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_OutMacErrorsTxPath extracts the value of the leaf OutMacErrorsTx from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_OutMacErrorsTxPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMacErrorsTx
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_OutMacPauseFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_OutMacPauseFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_OutMacPauseFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_OutMacPauseFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_OutMacPauseFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames to the batch object.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_OutMacPauseFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_OutMacPauseFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_OutMacPauseFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-mac-pause-frames to the batch object.
func (n *Interface_Ethernet_Counters_OutMacPauseFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_OutMacPauseFramesPath extracts the value of the leaf OutMacPauseFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_OutMacPauseFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMacPauseFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Ethernet_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_Counters_Out_8021QFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_Counters_Out_8021QFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a ONCE subscription.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_Out_8021QFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_Counters_Out_8021QFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_Out_8021QFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames to the batch object.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_Counters_Out_8021QFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Ethernet_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_Counters_Out_8021QFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Ethernet_Counters_Out_8021QFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/counters/out-8021q-frames to the batch object.
func (n *Interface_Ethernet_Counters_Out_8021QFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_Counters_Out_8021QFramesPath extracts the value of the leaf Out_8021QFrames from its parent oc.Interface_Ethernet_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Ethernet_Counters_Out_8021QFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Out_8021QFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_DuplexModePath) Lookup(t testing.TB) *oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_DuplexModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_DuplexModePath) Get(t testing.TB) oc.E_Ethernet_DuplexMode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_DuplexModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ethernet_DuplexMode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_DuplexModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a ONCE subscription.
func (n *Interface_Ethernet_DuplexModePathAny) Get(t testing.TB) []oc.E_Ethernet_DuplexMode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ethernet_DuplexMode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_DuplexModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ethernet_DuplexMode {
	t.Helper()
	c := &oc.CollectionE_Ethernet_DuplexMode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ethernet_DuplexMode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_DuplexModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ethernet_DuplexMode) bool) *oc.E_Ethernet_DuplexModeWatcher {
	t.Helper()
	w := &oc.E_Ethernet_DuplexModeWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_DuplexModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ethernet_DuplexMode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_DuplexModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ethernet_DuplexMode) bool) *oc.E_Ethernet_DuplexModeWatcher {
	t.Helper()
	return watch_Interface_Ethernet_DuplexModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_DuplexModePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ethernet_DuplexMode) *oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ethernet_DuplexMode) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode to the batch object.
func (n *Interface_Ethernet_DuplexModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_DuplexModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ethernet_DuplexMode {
	t.Helper()
	c := &oc.CollectionE_Ethernet_DuplexMode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ethernet_DuplexMode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_DuplexModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ethernet_DuplexMode) bool) *oc.E_Ethernet_DuplexModeWatcher {
	t.Helper()
	w := &oc.E_Ethernet_DuplexModeWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_DuplexModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ethernet_DuplexMode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_DuplexModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ethernet_DuplexMode) bool) *oc.E_Ethernet_DuplexModeWatcher {
	t.Helper()
	return watch_Interface_Ethernet_DuplexModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/duplex-mode to the batch object.
func (n *Interface_Ethernet_DuplexModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_DuplexModePath extracts the value of the leaf DuplexMode from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ethernet_DuplexMode.
func convertInterface_Ethernet_DuplexModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	qv := &oc.QualifiedE_Ethernet_DuplexMode{
		Metadata: md,
	}
	val := parent.DuplexMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_EnableFlowControlPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_EnableFlowControlPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableFlowControl())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_EnableFlowControlPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_EnableFlowControlPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a ONCE subscription.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_EnableFlowControlPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_EnableFlowControlPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_EnableFlowControlPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_EnableFlowControlPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Ethernet_EnableFlowControlPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_EnableFlowControlPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control to the batch object.
func (n *Interface_Ethernet_EnableFlowControlPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_EnableFlowControlPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_EnableFlowControlPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Ethernet_EnableFlowControlPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/enable-flow-control to the batch object.
func (n *Interface_Ethernet_EnableFlowControlPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_EnableFlowControlPath extracts the value of the leaf EnableFlowControl from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Ethernet_EnableFlowControlPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableFlowControl
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_FecModePath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_FecModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_FecModePath) Get(t testing.TB) oc.E_IfEthernet_INTERFACE_FEC {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_FecModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_INTERFACE_FEC
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_FecModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a ONCE subscription.
func (n *Interface_Ethernet_FecModePathAny) Get(t testing.TB) []oc.E_IfEthernet_INTERFACE_FEC {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_INTERFACE_FEC
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_FecModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_INTERFACE_FEC{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_FecModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool) *oc.E_IfEthernet_INTERFACE_FECWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_INTERFACE_FECWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_FecModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_INTERFACE_FEC)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_FecModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool) *oc.E_IfEthernet_INTERFACE_FECWatcher {
	t.Helper()
	return watch_Interface_Ethernet_FecModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_FecModePath) Await(t testing.TB, timeout time.Duration, val oc.E_IfEthernet_INTERFACE_FEC) *oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode to the batch object.
func (n *Interface_Ethernet_FecModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_FecModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_INTERFACE_FEC{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_FecModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool) *oc.E_IfEthernet_INTERFACE_FECWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_INTERFACE_FECWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_FecModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_INTERFACE_FEC)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_FecModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_INTERFACE_FEC) bool) *oc.E_IfEthernet_INTERFACE_FECWatcher {
	t.Helper()
	return watch_Interface_Ethernet_FecModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/fec-mode to the batch object.
func (n *Interface_Ethernet_FecModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_FecModePath extracts the value of the leaf FecMode from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_INTERFACE_FEC.
func convertInterface_Ethernet_FecModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_INTERFACE_FEC{
		Metadata: md,
	}
	val := parent.FecMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_HwMacAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_HwMacAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_HwMacAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_HwMacAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_HwMacAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a ONCE subscription.
func (n *Interface_Ethernet_HwMacAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_HwMacAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_HwMacAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_HwMacAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_HwMacAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ethernet_HwMacAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_HwMacAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address to the batch object.
func (n *Interface_Ethernet_HwMacAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_HwMacAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_HwMacAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_HwMacAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_HwMacAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ethernet_HwMacAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/hw-mac-address to the batch object.
func (n *Interface_Ethernet_HwMacAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_HwMacAddressPath extracts the value of the leaf HwMacAddress from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ethernet_HwMacAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.HwMacAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_MacAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_MacAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_MacAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_MacAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_MacAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a ONCE subscription.
func (n *Interface_Ethernet_MacAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_MacAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_MacAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_MacAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_MacAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ethernet_MacAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_MacAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address to the batch object.
func (n *Interface_Ethernet_MacAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_MacAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_MacAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_MacAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_MacAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ethernet_MacAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/mac-address to the batch object.
func (n *Interface_Ethernet_MacAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_MacAddressPath extracts the value of the leaf MacAddress from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ethernet_MacAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MacAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_NegotiatedDuplexModePath) Lookup(t testing.TB) *oc.QualifiedE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_NegotiatedDuplexModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_NegotiatedDuplexModePath) Get(t testing.TB) oc.E_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_NegotiatedDuplexModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ethernet_NegotiatedDuplexMode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_NegotiatedDuplexModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a ONCE subscription.
func (n *Interface_Ethernet_NegotiatedDuplexModePathAny) Get(t testing.TB) []oc.E_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ethernet_NegotiatedDuplexMode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_NegotiatedDuplexModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	c := &oc.CollectionE_Ethernet_NegotiatedDuplexMode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_NegotiatedDuplexModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool) *oc.E_Ethernet_NegotiatedDuplexModeWatcher {
	t.Helper()
	w := &oc.E_Ethernet_NegotiatedDuplexModeWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_NegotiatedDuplexModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ethernet_NegotiatedDuplexMode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_NegotiatedDuplexModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool) *oc.E_Ethernet_NegotiatedDuplexModeWatcher {
	t.Helper()
	return watch_Interface_Ethernet_NegotiatedDuplexModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_NegotiatedDuplexModePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ethernet_NegotiatedDuplexMode) *oc.QualifiedE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode to the batch object.
func (n *Interface_Ethernet_NegotiatedDuplexModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_NegotiatedDuplexModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	c := &oc.CollectionE_Ethernet_NegotiatedDuplexMode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_NegotiatedDuplexModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool) *oc.E_Ethernet_NegotiatedDuplexModeWatcher {
	t.Helper()
	w := &oc.E_Ethernet_NegotiatedDuplexModeWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_NegotiatedDuplexModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ethernet_NegotiatedDuplexMode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_NegotiatedDuplexModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ethernet_NegotiatedDuplexMode) bool) *oc.E_Ethernet_NegotiatedDuplexModeWatcher {
	t.Helper()
	return watch_Interface_Ethernet_NegotiatedDuplexModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-duplex-mode to the batch object.
func (n *Interface_Ethernet_NegotiatedDuplexModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_NegotiatedDuplexModePath extracts the value of the leaf NegotiatedDuplexMode from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ethernet_NegotiatedDuplexMode.
func convertInterface_Ethernet_NegotiatedDuplexModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	qv := &oc.QualifiedE_Ethernet_NegotiatedDuplexMode{
		Metadata: md,
	}
	val := parent.NegotiatedDuplexMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_NegotiatedPortSpeedPath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_NegotiatedPortSpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_NegotiatedPortSpeedPath) Get(t testing.TB) oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_NegotiatedPortSpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_NegotiatedPortSpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a ONCE subscription.
func (n *Interface_Ethernet_NegotiatedPortSpeedPathAny) Get(t testing.TB) []oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_ETHERNET_SPEED
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_NegotiatedPortSpeedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_ETHERNET_SPEED{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_NegotiatedPortSpeedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_ETHERNET_SPEEDWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_NegotiatedPortSpeedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_ETHERNET_SPEED)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_NegotiatedPortSpeedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	return watch_Interface_Ethernet_NegotiatedPortSpeedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_NegotiatedPortSpeedPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfEthernet_ETHERNET_SPEED) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed to the batch object.
func (n *Interface_Ethernet_NegotiatedPortSpeedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_NegotiatedPortSpeedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_ETHERNET_SPEED{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_NegotiatedPortSpeedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_ETHERNET_SPEEDWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_NegotiatedPortSpeedPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_ETHERNET_SPEED)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_NegotiatedPortSpeedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	return watch_Interface_Ethernet_NegotiatedPortSpeedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/negotiated-port-speed to the batch object.
func (n *Interface_Ethernet_NegotiatedPortSpeedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_NegotiatedPortSpeedPath extracts the value of the leaf NegotiatedPortSpeed from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_ETHERNET_SPEED.
func convertInterface_Ethernet_NegotiatedPortSpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_ETHERNET_SPEED{
		Metadata: md,
	}
	val := parent.NegotiatedPortSpeed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_PortSpeedPath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_PortSpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_PortSpeedPath) Get(t testing.TB) oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_PortSpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_PortSpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a ONCE subscription.
func (n *Interface_Ethernet_PortSpeedPathAny) Get(t testing.TB) []oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_ETHERNET_SPEED
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_PortSpeedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_ETHERNET_SPEED{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_PortSpeedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_ETHERNET_SPEEDWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_PortSpeedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_ETHERNET_SPEED)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_PortSpeedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	return watch_Interface_Ethernet_PortSpeedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_PortSpeedPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfEthernet_ETHERNET_SPEED) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed to the batch object.
func (n *Interface_Ethernet_PortSpeedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_PortSpeedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_ETHERNET_SPEED{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_PortSpeedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_ETHERNET_SPEEDWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_PortSpeedPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_ETHERNET_SPEED)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_PortSpeedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	return watch_Interface_Ethernet_PortSpeedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/port-speed to the batch object.
func (n *Interface_Ethernet_PortSpeedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_PortSpeedPath extracts the value of the leaf PortSpeed from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_ETHERNET_SPEED.
func convertInterface_Ethernet_PortSpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_ETHERNET_SPEED{
		Metadata: md,
	}
	val := parent.PortSpeed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Ethernet{}
	md, ok := oc.Lookup(t, n, "Interface_Ethernet", goStruct, true, false)
	if ok {
		return convertInterface_Ethernet_StandaloneLinkTrainingPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetStandaloneLinkTraining())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ethernet_StandaloneLinkTrainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a ONCE subscription.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_StandaloneLinkTrainingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ethernet_StandaloneLinkTrainingPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Ethernet_StandaloneLinkTrainingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training to the batch object.
func (n *Interface_Ethernet_StandaloneLinkTrainingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ethernet_StandaloneLinkTrainingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ethernet", structs[pre], queryPath, true, false)
			qv := convertInterface_Ethernet_StandaloneLinkTrainingPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Ethernet_StandaloneLinkTrainingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/ethernet/state/standalone-link-training to the batch object.
func (n *Interface_Ethernet_StandaloneLinkTrainingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ethernet_StandaloneLinkTrainingPath extracts the value of the leaf StandaloneLinkTraining from its parent oc.Interface_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Ethernet_StandaloneLinkTrainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ethernet) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.StandaloneLinkTraining
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
