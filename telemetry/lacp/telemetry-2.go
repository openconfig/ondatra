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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_CountersPath) Lookup(t testing.TB) *oc.QualifiedLacp_Interface_Member_Counters {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp_Interface_Member_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_CountersPath) Get(t testing.TB) *oc.Lacp_Interface_Member_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp_Interface_Member_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp_Interface_Member_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp_Interface_Member_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a ONCE subscription.
func (n *Lacp_Interface_Member_CountersPathAny) Get(t testing.TB) []*oc.Lacp_Interface_Member_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp_Interface_Member_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_Interface_Member_Counters {
	t.Helper()
	c := &oc.CollectionLacp_Interface_Member_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_Interface_Member_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp_Interface_Member_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp_Interface_Member_Counters)))
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_Interface_Member_Counters) bool) *oc.Lacp_Interface_Member_CountersWatcher {
	t.Helper()
	w := &oc.Lacp_Interface_Member_CountersWatcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedLacp_Interface_Member_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_Interface_Member_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_Interface_Member_Counters) bool) *oc.Lacp_Interface_Member_CountersWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp_Interface_Member_Counters) *oc.QualifiedLacp_Interface_Member_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp_Interface_Member_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters to the batch object.
func (n *Lacp_Interface_Member_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_Interface_Member_Counters {
	t.Helper()
	c := &oc.CollectionLacp_Interface_Member_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_Interface_Member_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_Interface_Member_Counters) bool) *oc.Lacp_Interface_Member_CountersWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_CountersPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters to the batch object.
func (n *Lacp_Interface_Member_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLacp_Interface_Member_Counters_LacpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return convertLacp_Interface_Member_Counters_LacpErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-errors to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpErrorsPath extracts the value of the leaf LacpErrors from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpInPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLacp_Interface_Member_Counters_LacpInPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpInPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return convertLacp_Interface_Member_Counters_LacpInPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpInPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpInPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpInPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpInPktsPath extracts the value of the leaf LacpInPkts from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpInPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpInPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member_Counters", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_Counters_LacpOutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLacp_Interface_Member_Counters_LacpOutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a ONCE subscription.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_Counters_LacpOutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_Interface_Member_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member_Counters", gs, queryPath, true, false)
		return convertLacp_Interface_Member_Counters_LacpOutPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpOutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_Interface_Member_Counters_LacpOutPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts to the batch object.
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_Counters_LacpOutPktsPath extracts the value of the leaf LacpOutPkts from its parent oc.Lacp_Interface_Member_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_Interface_Member_Counters_LacpOutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpOutPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
