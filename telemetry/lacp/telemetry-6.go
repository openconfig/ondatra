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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_SynchronizationPath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpSynchronizationType {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_SynchronizationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_SynchronizationPath) Get(t testing.TB) oc.E_Lacp_LacpSynchronizationType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_SynchronizationPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpSynchronizationType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpSynchronizationType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_SynchronizationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a ONCE subscription.
func (n *Lacp_Interface_Member_SynchronizationPathAny) Get(t testing.TB) []oc.E_Lacp_LacpSynchronizationType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpSynchronizationType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_SynchronizationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpSynchronizationType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpSynchronizationType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpSynchronizationType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_SynchronizationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpSynchronizationType) bool) *oc.E_Lacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpSynchronizationTypeWatcher{}
	gs := &oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_SynchronizationPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpSynchronizationType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_SynchronizationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpSynchronizationType) bool) *oc.E_Lacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_SynchronizationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_SynchronizationPath) Await(t testing.TB, timeout time.Duration, val oc.E_Lacp_LacpSynchronizationType) *oc.QualifiedE_Lacp_LacpSynchronizationType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lacp_LacpSynchronizationType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization to the batch object.
func (n *Lacp_Interface_Member_SynchronizationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_SynchronizationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpSynchronizationType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpSynchronizationType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpSynchronizationType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_SynchronizationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpSynchronizationType) bool) *oc.E_Lacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpSynchronizationTypeWatcher{}
	structs := map[string]*oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_SynchronizationPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpSynchronizationType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_SynchronizationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpSynchronizationType) bool) *oc.E_Lacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_SynchronizationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization to the batch object.
func (n *Lacp_Interface_Member_SynchronizationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_SynchronizationPath extracts the value of the leaf Synchronization from its parent oc.Lacp_Interface_Member
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpSynchronizationType.
func convertLacp_Interface_Member_SynchronizationPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member) *oc.QualifiedE_Lacp_LacpSynchronizationType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpSynchronizationType{
		Metadata: md,
	}
	val := parent.Synchronization
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_SystemIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_SystemIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_SystemIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_SystemIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_SystemIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a ONCE subscription.
func (n *Lacp_Interface_Member_SystemIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_SystemIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_SystemIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_SystemIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_SystemIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_SystemIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_SystemIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id to the batch object.
func (n *Lacp_Interface_Member_SystemIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_SystemIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_SystemIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_SystemIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_SystemIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_SystemIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/system-id to the batch object.
func (n *Lacp_Interface_Member_SystemIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_SystemIdPath extracts the value of the leaf SystemId from its parent oc.Lacp_Interface_Member
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLacp_Interface_Member_SystemIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpTimeoutType {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_TimeoutPath) Get(t testing.TB) oc.E_Lacp_LacpTimeoutType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpTimeoutType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpTimeoutType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a ONCE subscription.
func (n *Lacp_Interface_Member_TimeoutPathAny) Get(t testing.TB) []oc.E_Lacp_LacpTimeoutType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpTimeoutType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_TimeoutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpTimeoutType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpTimeoutType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpTimeoutType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_TimeoutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpTimeoutType) bool) *oc.E_Lacp_LacpTimeoutTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpTimeoutTypeWatcher{}
	gs := &oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_Member_TimeoutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpTimeoutType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_TimeoutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpTimeoutType) bool) *oc.E_Lacp_LacpTimeoutTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_TimeoutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_TimeoutPath) Await(t testing.TB, timeout time.Duration, val oc.E_Lacp_LacpTimeoutType) *oc.QualifiedE_Lacp_LacpTimeoutType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lacp_LacpTimeoutType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout to the batch object.
func (n *Lacp_Interface_Member_TimeoutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_TimeoutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpTimeoutType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpTimeoutType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpTimeoutType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_TimeoutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpTimeoutType) bool) *oc.E_Lacp_LacpTimeoutTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpTimeoutTypeWatcher{}
	structs := map[string]*oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface_Member{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface_Member", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_Member_TimeoutPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpTimeoutType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_TimeoutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpTimeoutType) bool) *oc.E_Lacp_LacpTimeoutTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_TimeoutPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/timeout to the batch object.
func (n *Lacp_Interface_Member_TimeoutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_TimeoutPath extracts the value of the leaf Timeout from its parent oc.Lacp_Interface_Member
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpTimeoutType.
func convertLacp_Interface_Member_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member) *oc.QualifiedE_Lacp_LacpTimeoutType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpTimeoutType{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, false)
	if ok {
		return convertLacp_Interface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/name with a ONCE subscription.
func (n *Lacp_Interface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_Interface_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/name to the batch object.
func (n *Lacp_Interface_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_Interface_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/name to the batch object.
func (n *Lacp_Interface_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_NamePath extracts the value of the leaf Name from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLacp_Interface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedString {
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
