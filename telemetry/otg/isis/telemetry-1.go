package isis

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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_OutLspPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_OutLspPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_OutLspPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_OutLspPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_OutLspPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_OutLspPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutLspPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutLspPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_OutLspPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutLspPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutLspPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_OutLspPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp to the batch object.
func (n *IsisRouter_Counters_Level1_OutLspPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutLspPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutLspPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_OutLspPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutLspPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutLspPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-lsp to the batch object.
func (n *IsisRouter_Counters_Level1_OutLspPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_OutLspPath extracts the value of the leaf OutLsp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_OutLspPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutLsp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_OutP2PHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_OutP2PHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutP2PHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_OutP2PHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutP2PHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutP2PHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_OutP2PHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutP2PHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_OutP2PHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_OutP2PHellosPath extracts the value of the leaf OutP2PHellos from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_OutP2PHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutP2PHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_OutPsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_OutPsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_OutPsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_OutPsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_OutPsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_OutPsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutPsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutPsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_OutPsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutPsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutPsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_OutPsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp to the batch object.
func (n *IsisRouter_Counters_Level1_OutPsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutPsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutPsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_OutPsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutPsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutPsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-psnp to the batch object.
func (n *IsisRouter_Counters_Level1_OutPsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_OutPsnpPath extracts the value of the leaf OutPsnp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_OutPsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPsnp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_SessionsFlapPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_SessionsFlapPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_SessionsFlapPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_SessionsFlapPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_SessionsFlapPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_SessionsFlapPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_SessionsFlapPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_SessionsFlapPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_SessionsFlapPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_SessionsFlapPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_SessionsFlapPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_SessionsFlapPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap to the batch object.
func (n *IsisRouter_Counters_Level1_SessionsFlapPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_SessionsFlapPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_SessionsFlapPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_SessionsFlapPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_SessionsFlapPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_SessionsFlapPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-flap to the batch object.
func (n *IsisRouter_Counters_Level1_SessionsFlapPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_SessionsFlapPath extracts the value of the leaf SessionsFlap from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_SessionsFlapPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.SessionsFlap
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_SessionsUpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_SessionsUpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_SessionsUpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_SessionsUpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_SessionsUpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_SessionsUpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_SessionsUpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_SessionsUpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_SessionsUpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_SessionsUpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_SessionsUpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_SessionsUpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up to the batch object.
func (n *IsisRouter_Counters_Level1_SessionsUpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_SessionsUpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_SessionsUpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_SessionsUpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_SessionsUpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_SessionsUpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/sessions-up to the batch object.
func (n *IsisRouter_Counters_Level1_SessionsUpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_SessionsUpPath extracts the value of the leaf SessionsUp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_SessionsUpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.SessionsUp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2Path) Lookup(t testing.TB) *oc.QualifiedIsisRouter_Counters_Level2 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_Counters_Level2{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2Path) Get(t testing.TB) *oc.IsisRouter_Counters_Level2 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2PathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_Counters_Level2 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_Counters_Level2
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_Counters_Level2{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a ONCE subscription.
func (n *IsisRouter_Counters_Level2PathAny) Get(t testing.TB) []*oc.IsisRouter_Counters_Level2 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_Counters_Level2
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_Counters_Level2 {
	t.Helper()
	c := &oc.CollectionIsisRouter_Counters_Level2{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_Counters_Level2) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_Counters_Level2{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_Counters_Level2)))
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level2) bool) *oc.IsisRouter_Counters_Level2Watcher {
	t.Helper()
	w := &oc.IsisRouter_Counters_Level2Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_Counters_Level2{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_Counters_Level2)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level2) bool) *oc.IsisRouter_Counters_Level2Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2Path(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2Path) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_Counters_Level2) *oc.QualifiedIsisRouter_Counters_Level2 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_Counters_Level2) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 to the batch object.
func (n *IsisRouter_Counters_Level2Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_Counters_Level2 {
	t.Helper()
	c := &oc.CollectionIsisRouter_Counters_Level2{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_Counters_Level2) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level2) bool) *oc.IsisRouter_Counters_Level2Watcher {
	t.Helper()
	w := &oc.IsisRouter_Counters_Level2Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_Counters_Level2{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_Counters_Level2)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level2) bool) *oc.IsisRouter_Counters_Level2Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2PathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2 to the batch object.
func (n *IsisRouter_Counters_Level2PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_DatabaseSizePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_DatabaseSizePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_DatabaseSizePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_DatabaseSizePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_DatabaseSizePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_DatabaseSizePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_DatabaseSizePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_DatabaseSizePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_DatabaseSizePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_DatabaseSizePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_DatabaseSizePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_DatabaseSizePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size to the batch object.
func (n *IsisRouter_Counters_Level2_DatabaseSizePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_DatabaseSizePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_DatabaseSizePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_DatabaseSizePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_DatabaseSizePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_DatabaseSizePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/database-size to the batch object.
func (n *IsisRouter_Counters_Level2_DatabaseSizePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_DatabaseSizePath extracts the value of the leaf DatabaseSize from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_DatabaseSizePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DatabaseSize
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_InBcastHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_InBcastHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_InBcastHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_InBcastHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_InBcastHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_InBcastHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InBcastHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InBcastHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_InBcastHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InBcastHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InBcastHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_InBcastHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_InBcastHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InBcastHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InBcastHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_InBcastHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InBcastHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InBcastHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_InBcastHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_InBcastHellosPath extracts the value of the leaf InBcastHellos from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_InBcastHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InBcastHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_InCsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_InCsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_InCsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_InCsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_InCsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_InCsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InCsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InCsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_InCsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InCsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InCsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_InCsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp to the batch object.
func (n *IsisRouter_Counters_Level2_InCsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InCsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InCsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_InCsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InCsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InCsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-csnp to the batch object.
func (n *IsisRouter_Counters_Level2_InCsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_InCsnpPath extracts the value of the leaf InCsnp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_InCsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InCsnp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_InLspPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_InLspPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_InLspPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_InLspPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_InLspPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_InLspPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InLspPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InLspPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_InLspPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InLspPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InLspPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_InLspPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp to the batch object.
func (n *IsisRouter_Counters_Level2_InLspPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InLspPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InLspPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_InLspPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InLspPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InLspPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-lsp to the batch object.
func (n *IsisRouter_Counters_Level2_InLspPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_InLspPath extracts the value of the leaf InLsp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_InLspPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InLsp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_InP2PHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_InP2PHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_InP2PHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_InP2PHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_InP2PHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_InP2PHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InP2PHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InP2PHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_InP2PHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InP2PHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InP2PHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_InP2PHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_InP2PHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InP2PHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InP2PHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_InP2PHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InP2PHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InP2PHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_InP2PHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_InP2PHellosPath extracts the value of the leaf InP2PHellos from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_InP2PHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InP2PHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
