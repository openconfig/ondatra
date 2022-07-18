package lacp

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

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_PartnerPortNumPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_PartnerPortNumPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_PartnerPortNumPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_PartnerPortNumPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_PartnerPortNumPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a ONCE subscription.
func (n *Lacp_LagMember_PartnerPortNumPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_PartnerPortNumPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_PartnerPortNumPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_PartnerPortNumPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_PartnerPortNumPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_LagMember_PartnerPortNumPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_PartnerPortNumPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num to the batch object.
func (n *Lacp_LagMember_PartnerPortNumPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_PartnerPortNumPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_PartnerPortNumPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_PartnerPortNumPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_PartnerPortNumPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_LagMember_PartnerPortNumPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/partner-port-num to the batch object.
func (n *Lacp_LagMember_PartnerPortNumPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_PartnerPortNumPath extracts the value of the leaf PartnerPortNum from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertLacp_LagMember_PartnerPortNumPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PartnerPortNum
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_PortNumPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_PortNumPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_PortNumPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_PortNumPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_PortNumPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a ONCE subscription.
func (n *Lacp_LagMember_PortNumPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_PortNumPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_PortNumPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_PortNumPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_PortNumPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_LagMember_PortNumPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_PortNumPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num to the batch object.
func (n *Lacp_LagMember_PortNumPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_PortNumPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_PortNumPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_PortNumPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_PortNumPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_LagMember_PortNumPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/port-num to the batch object.
func (n *Lacp_LagMember_PortNumPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_PortNumPath extracts the value of the leaf PortNum from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertLacp_LagMember_PortNumPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PortNum
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_SynchronizationPath) Lookup(t testing.TB) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_SynchronizationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_SynchronizationPath) Get(t testing.TB) oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_SynchronizationPathAny) Lookup(t testing.TB) []*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_SynchronizationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a ONCE subscription.
func (n *Lacp_LagMember_SynchronizationPathAny) Get(t testing.TB) []oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_SynchronizationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	c := &oc.CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_SynchronizationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	w := &oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_SynchronizationPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_SynchronizationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	return watch_Lacp_LagMember_SynchronizationPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_SynchronizationPath) Await(t testing.TB, timeout time.Duration, val oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationType) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization to the batch object.
func (n *Lacp_LagMember_SynchronizationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_SynchronizationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	c := &oc.CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_SynchronizationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	w := &oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_SynchronizationPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_SynchronizationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher {
	t.Helper()
	return watch_Lacp_LagMember_SynchronizationPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/synchronization to the batch object.
func (n *Lacp_LagMember_SynchronizationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_SynchronizationPath extracts the value of the leaf Synchronization from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType.
func convertLacp_LagMember_SynchronizationPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	qv := &oc.QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType{
		Metadata: md,
	}
	val := parent.Synchronization
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_SystemIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_SystemIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_SystemIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_SystemIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_SystemIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a ONCE subscription.
func (n *Lacp_LagMember_SystemIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_SystemIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_SystemIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_SystemIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_SystemIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_LagMember_SystemIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_SystemIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id to the batch object.
func (n *Lacp_LagMember_SystemIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_SystemIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_SystemIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_SystemIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_SystemIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_LagMember_SystemIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/system-id to the batch object.
func (n *Lacp_LagMember_SystemIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_SystemIdPath extracts the value of the leaf SystemId from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLacp_LagMember_SystemIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedString {
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

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_TimeoutPath) Get(t testing.TB) oc.E_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a ONCE subscription.
func (n *Lacp_LagMember_TimeoutPathAny) Get(t testing.TB) []oc.E_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_OpenTrafficGeneratorLacp_LacpTimeoutType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_TimeoutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	c := &oc.CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_TimeoutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher {
	t.Helper()
	w := &oc.E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_TimeoutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_TimeoutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher {
	t.Helper()
	return watch_Lacp_LagMember_TimeoutPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_TimeoutPath) Await(t testing.TB, timeout time.Duration, val oc.E_OpenTrafficGeneratorLacp_LacpTimeoutType) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout to the batch object.
func (n *Lacp_LagMember_TimeoutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_TimeoutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	c := &oc.CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_TimeoutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher {
	t.Helper()
	w := &oc.E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_TimeoutPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_TimeoutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher {
	t.Helper()
	return watch_Lacp_LagMember_TimeoutPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/timeout to the batch object.
func (n *Lacp_LagMember_TimeoutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_TimeoutPath extracts the value of the leaf Timeout from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType.
func convertLacp_LagMember_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	qv := &oc.QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
