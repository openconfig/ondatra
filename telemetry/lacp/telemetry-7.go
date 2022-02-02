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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_SystemIdMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, false)
	if ok {
		return convertLacp_Interface_SystemIdMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_SystemIdMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_SystemIdMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLacp_Interface_SystemIdMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a ONCE subscription.
func (n *Lacp_Interface_SystemIdMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_SystemIdMacPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_SystemIdMacPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface", gs, queryPath, true, false)
		return convertLacp_Interface_SystemIdMacPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_SystemIdMacPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_Interface_SystemIdMacPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_SystemIdMacPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac to the batch object.
func (n *Lacp_Interface_SystemIdMacPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_SystemIdMacPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_SystemIdMacPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lacp_Interface_SystemIdMacPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac to the batch object.
func (n *Lacp_Interface_SystemIdMacPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_SystemIdMacPath extracts the value of the leaf SystemIdMac from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLacp_Interface_SystemIdMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemIdMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_SystemPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, false)
	if ok {
		return convertLacp_Interface_SystemPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_SystemPriorityPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_SystemPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_SystemPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a ONCE subscription.
func (n *Lacp_Interface_SystemPriorityPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_SystemPriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_SystemPriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface", gs, queryPath, true, false)
		return convertLacp_Interface_SystemPriorityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_SystemPriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_Interface_SystemPriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_SystemPriorityPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/state/system-priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/system-priority to the batch object.
func (n *Lacp_Interface_SystemPriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_SystemPriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/system-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_SystemPriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_Interface_SystemPriorityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/system-priority to the batch object.
func (n *Lacp_Interface_SystemPriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_SystemPriorityPath extracts the value of the leaf SystemPriority from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertLacp_Interface_SystemPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SystemPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/state/system-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_SystemPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Lacp{}
	md, ok := oc.Lookup(t, n, "Lacp", goStruct, true, false)
	if ok {
		return convertLacp_SystemPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/state/system-priority with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_SystemPriorityPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/state/system-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_SystemPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_SystemPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/state/system-priority with a ONCE subscription.
func (n *Lacp_SystemPriorityPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/state/system-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_SystemPriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_SystemPriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Lacp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp", gs, queryPath, true, false)
		return convertLacp_SystemPriorityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/state/system-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_SystemPriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_SystemPriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/state/system-priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_SystemPriorityPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/state/system-priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/state/system-priority to the batch object.
func (n *Lacp_SystemPriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/state/system-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_SystemPriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/state/system-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_SystemPriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Lacp_SystemPriorityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/state/system-priority to the batch object.
func (n *Lacp_SystemPriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_SystemPriorityPath extracts the value of the leaf SystemPriority from its parent oc.Lacp
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertLacp_SystemPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SystemPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
