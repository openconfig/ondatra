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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_SystemNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_SystemNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_SystemNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_SystemNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_SystemNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_SystemNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_SystemNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_SystemNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_SystemNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_SystemNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_SystemNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_SystemNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name to the batch object.
func (n *Lldp_Interface_Neighbor_SystemNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_SystemNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_SystemNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertLldp_Interface_Neighbor_SystemNamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_SystemNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_SystemNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/system-name to the batch object.
func (n *Lldp_Interface_Neighbor_SystemNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_SystemNamePath extracts the value of the leaf SystemName from its parent oc.Lldp_Interface_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_Neighbor_SystemNamePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_TlvPath) Lookup(t testing.TB) *oc.QualifiedLldp_Interface_Neighbor_Tlv {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Tlv", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp_Interface_Neighbor_Tlv{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_TlvPath) Get(t testing.TB) *oc.Lldp_Interface_Neighbor_Tlv {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_TlvPathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Interface_Neighbor_Tlv {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Interface_Neighbor_Tlv
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Interface_Neighbor_Tlv{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_TlvPathAny) Get(t testing.TB) []*oc.Lldp_Interface_Neighbor_Tlv {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Interface_Neighbor_Tlv
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_TlvPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Neighbor_Tlv {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Neighbor_Tlv{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp_Interface_Neighbor_Tlv{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp_Interface_Neighbor_Tlv)))
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_TlvPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool) *oc.Lldp_Interface_Neighbor_TlvWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_Neighbor_TlvWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldp_Interface_Neighbor_Tlv{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Neighbor_Tlv)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_TlvPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool) *oc.Lldp_Interface_Neighbor_TlvWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_TlvPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_TlvPath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp_Interface_Neighbor_Tlv) *oc.QualifiedLldp_Interface_Neighbor_Tlv {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv to the batch object.
func (n *Lldp_Interface_Neighbor_TlvPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_TlvPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Interface_Neighbor_Tlv {
	t.Helper()
	c := &oc.CollectionLldp_Interface_Neighbor_Tlv{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_TlvPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool) *oc.Lldp_Interface_Neighbor_TlvWatcher {
	t.Helper()
	w := &oc.Lldp_Interface_Neighbor_TlvWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldp_Interface_Neighbor_Tlv{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Interface_Neighbor_Tlv)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_TlvPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Interface_Neighbor_Tlv) bool) *oc.Lldp_Interface_Neighbor_TlvWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_TlvPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv to the batch object.
func (n *Lldp_Interface_Neighbor_TlvPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_Tlv_OuiPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_Tlv_OuiPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_OuiPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_Tlv_OuiPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_OuiPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_OuiPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_Tlv_OuiPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_OuiPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_OuiPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_Tlv_OuiPath extracts the value of the leaf Oui from its parent oc.Lldp_Interface_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_Neighbor_Tlv_OuiPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor_Tlv) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Oui
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_Tlv_OuiSubtypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_Tlv_OuiSubtypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_OuiSubtypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_Tlv_OuiSubtypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_OuiSubtypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_Tlv_OuiSubtypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_OuiSubtypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_Tlv_OuiSubtypePath extracts the value of the leaf OuiSubtype from its parent oc.Lldp_Interface_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_Neighbor_Tlv_OuiSubtypePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor_Tlv) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OuiSubtype
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_TypePath) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_Tlv_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_Tlv_TypePath) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_Tlv_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_Tlv_TypePathAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	gs := &oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_Tlv_TypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_Tlv_TypePath) Await(t testing.TB, timeout time.Duration, val int32) *oc.QualifiedInt32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_TypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_Tlv_TypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_TypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/type to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_Tlv_TypePath extracts the value of the leaf Type from its parent oc.Lldp_Interface_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertLldp_Interface_Neighbor_Tlv_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor_Tlv) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePath) Lookup(t testing.TB) *oc.QualifiedBinary {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_Tlv_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePath) Get(t testing.TB) oc.Binary {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedBinary {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBinary
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_Tlv_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePathAny) Get(t testing.TB) []oc.Binary {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBinary {
	t.Helper()
	c := &oc.CollectionBinary{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBinary) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_ValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	w := &oc.BinaryWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_Tlv_ValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBinary)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_ValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePath) Await(t testing.TB, timeout time.Duration, val oc.Binary) *oc.QualifiedBinary {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBinary) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBinary {
	t.Helper()
	c := &oc.CollectionBinary{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBinary) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Tlv_ValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	w := &oc.BinaryWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_Tlv_ValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBinary)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Tlv_ValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/custom-tlvs/tlv/state/value to the batch object.
func (n *Lldp_Interface_Neighbor_Tlv_ValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_Tlv_ValuePath extracts the value of the leaf Value from its parent oc.Lldp_Interface_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedBinary.
func convertLldp_Interface_Neighbor_Tlv_ValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor_Tlv) *oc.QualifiedBinary {
	t.Helper()
	qv := &oc.QualifiedBinary{
		Metadata: md,
	}
	val := parent.Value
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
