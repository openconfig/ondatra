package lacp

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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpRxErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_Counters_LacpRxErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpRxErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_Counters_LacpRxErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpRxErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpRxErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member_Counters", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_Counters_LacpRxErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpRxErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpRxErrorsPath extracts the value of the leaf LacpRxErrors from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpRxErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpRxErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member_Counters", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpTimeoutTransitionsPath extracts the value of the leaf LacpTimeoutTransitions from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpTimeoutTransitionsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpTimeoutTransitions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpTxErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_Counters_LacpTxErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpTxErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_Counters_LacpTxErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpTxErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpTxErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member_Counters", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_Counters_LacpTxErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpTxErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpTxErrorsPath extracts the value of the leaf LacpTxErrors from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpTxErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpTxErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpUnknownErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_Counters_LacpUnknownErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpUnknownErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_Counters_LacpUnknownErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpUnknownErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member_Counters", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_Counters_LacpUnknownErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpUnknownErrorsPath extracts the value of the leaf LacpUnknownErrors from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpUnknownErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpUnknownErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
