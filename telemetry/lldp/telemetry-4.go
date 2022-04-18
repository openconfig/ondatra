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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Counters_TlvUnknownPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Counters_TlvUnknownPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Counters_TlvUnknownPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Counters_TlvUnknownPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertLldp_Interface_Counters_TlvUnknownPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a ONCE subscription.
func (n *Lldp_Interface_Counters_TlvUnknownPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_TlvUnknownPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_TlvUnknownPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Counters_TlvUnknownPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_TlvUnknownPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_TlvUnknownPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Counters_TlvUnknownPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown to the batch object.
func (n *Lldp_Interface_Counters_TlvUnknownPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Counters_TlvUnknownPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Counters_TlvUnknownPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertLldp_Interface_Counters_TlvUnknownPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Counters_TlvUnknownPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Counters_TlvUnknownPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown to the batch object.
func (n *Lldp_Interface_Counters_TlvUnknownPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Counters_TlvUnknownPath extracts the value of the leaf TlvUnknown from its parent oc.Lldp_Interface_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Counters_TlvUnknownPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Lldp_Interface{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface", goStruct, true, false)
	if ok {
		return convertLldp_Interface_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a ONCE subscription.
func (n *Lldp_Interface_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Lldp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lldp_Interface_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/enabled to the batch object.
func (n *Lldp_Interface_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Lldp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_EnabledPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lldp_Interface_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/enabled to the batch object.
func (n *Lldp_Interface_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_EnabledPath extracts the value of the leaf Enabled from its parent oc.Lldp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLldp_Interface_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface", goStruct, true, false)
	if ok {
		return convertLldp_Interface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/state/name with a ONCE subscription.
func (n *Lldp_Interface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/name to the batch object.
func (n *Lldp_Interface_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/state/name to the batch object.
func (n *Lldp_Interface_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_NamePath extracts the value of the leaf Name from its parent oc.Lldp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_NeighborPath) Lookup(t testing.TB) *oc.QualifiedLldp_Interface_Neighbor {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp_Interface_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_NeighborPath) Get(t testing.TB) *oc.Lldp_Interface_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Interface_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Interface_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Interface_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a ONCE subscription.
func (n *Lldp_Interface_NeighborPathAny) Get(t testing.TB) []*oc.Lldp_Interface_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Interface_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Neighbor {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp_Interface_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp_Interface_Neighbor)))
		return false
	})
	return c
}

func watch_Lldp_Interface_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor) bool) *oc.Lldp_Interface_NeighborWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_NeighborWatcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldp_Interface_Neighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor) bool) *oc.Lldp_Interface_NeighborWatcher {
	t.Helper()
	return watch_Lldp_Interface_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp_Interface_Neighbor) *oc.QualifiedLldp_Interface_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp_Interface_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor to the batch object.
func (n *Lldp_Interface_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Neighbor {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_NeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor) bool) *oc.Lldp_Interface_NeighborWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_NeighborWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldp_Interface_Neighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor) bool) *oc.Lldp_Interface_NeighborWatcher {
	t.Helper()
	return watch_Lldp_Interface_NeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor to the batch object.
func (n *Lldp_Interface_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_AgePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_AgePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_AgePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_AgePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_AgePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_AgePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_AgePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_AgePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_AgePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_AgePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_AgePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_AgePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age to the batch object.
func (n *Lldp_Interface_Neighbor_AgePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_AgePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_AgePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_AgePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_AgePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_AgePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age to the batch object.
func (n *Lldp_Interface_Neighbor_AgePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_AgePath extracts the value of the leaf Age from its parent oc.Lldp_Interface_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Interface_Neighbor_AgePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Age
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_CapabilityPath) Lookup(t testing.TB) *oc.QualifiedLldp_Interface_Neighbor_Capability {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Capability{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Capability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp_Interface_Neighbor_Capability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_CapabilityPath) Get(t testing.TB) *oc.Lldp_Interface_Neighbor_Capability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_CapabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Interface_Neighbor_Capability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Interface_Neighbor_Capability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Capability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Interface_Neighbor_Capability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_CapabilityPathAny) Get(t testing.TB) []*oc.Lldp_Interface_Neighbor_Capability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Interface_Neighbor_Capability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_CapabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Neighbor_Capability {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Neighbor_Capability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Neighbor_Capability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp_Interface_Neighbor_Capability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp_Interface_Neighbor_Capability)))
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_CapabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Capability) bool) *oc.Lldp_Interface_Neighbor_CapabilityWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_Neighbor_CapabilityWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Capability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldp_Interface_Neighbor_Capability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Neighbor_Capability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_CapabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Capability) bool) *oc.Lldp_Interface_Neighbor_CapabilityWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_CapabilityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_CapabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp_Interface_Neighbor_Capability) *oc.QualifiedLldp_Interface_Neighbor_Capability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp_Interface_Neighbor_Capability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability to the batch object.
func (n *Lldp_Interface_Neighbor_CapabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_CapabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Neighbor_Capability {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Neighbor_Capability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Neighbor_Capability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_CapabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Capability) bool) *oc.Lldp_Interface_Neighbor_CapabilityWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_Neighbor_CapabilityWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Capability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Capability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldp_Interface_Neighbor_Capability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Neighbor_Capability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_CapabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Capability) bool) *oc.Lldp_Interface_Neighbor_CapabilityWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_CapabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability to the batch object.
func (n *Lldp_Interface_Neighbor_CapabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
