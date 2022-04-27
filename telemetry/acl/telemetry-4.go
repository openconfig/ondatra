package acl

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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_DestinationMacMaskPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_L2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask to the batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_L2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac-mask to the batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacMaskPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath extracts the value of the leaf DestinationMacMask from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_DestinationMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_DestinationMacPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_L2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_DestinationMacPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac to the batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_DestinationMacPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_L2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_DestinationMacPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/destination-mac to the batch object.
func (n *Acl_AclSet_AclEntry_L2_DestinationMacPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_L2_DestinationMacPath extracts the value of the leaf DestinationMac from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_DestinationMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_EthertypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_EthertypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_L2_Ethertype_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_EthertypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool) *oc.Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_L2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_L2_EthertypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool) *oc.Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_EthertypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Await(t testing.TB, timeout time.Duration, val oc.Acl_AclSet_AclEntry_L2_Ethertype_Union) *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype to the batch object.
func (n *Acl_AclSet_AclEntry_L2_EthertypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_EthertypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool) *oc.Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_L2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_L2_EthertypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) bool) *oc.Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_EthertypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/ethertype to the batch object.
func (n *Acl_AclSet_AclEntry_L2_EthertypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_L2_EthertypePath extracts the value of the leaf Ethertype from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union.
func convertAcl_AclSet_AclEntry_L2_EthertypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union{
		Metadata: md,
	}
	val := parent.Ethertype
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_SourceMacMaskPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_L2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_SourceMacMaskPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask to the batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_L2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac-mask to the batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacMaskPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath extracts the value of the leaf SourceMacMask from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_SourceMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_L2{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_L2", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_L2_SourceMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_L2_SourceMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_SourceMacPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_L2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_L2_SourceMacPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_SourceMacPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac to the batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_L2_SourceMacPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_L2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_L2", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_L2_SourceMacPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_L2_SourceMacPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/l2/state/source-mac to the batch object.
func (n *Acl_AclSet_AclEntry_L2_SourceMacPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_L2_SourceMacPath extracts the value of the leaf SourceMac from its parent oc.Acl_AclSet_AclEntry_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_L2_SourceMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_MatchedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_MatchedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_MatchedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_MatchedOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_MatchedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets to the batch object.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_MatchedOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_MatchedOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_MatchedOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-octets to the batch object.
func (n *Acl_AclSet_AclEntry_MatchedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_MatchedOctetsPath extracts the value of the leaf MatchedOctets from its parent oc.Acl_AclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertAcl_AclSet_AclEntry_MatchedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_MatchedPacketsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_MatchedPacketsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_MatchedPacketsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_MatchedPacketsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_MatchedPacketsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets to the batch object.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_MatchedPacketsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_MatchedPacketsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_MatchedPacketsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/matched-packets to the batch object.
func (n *Acl_AclSet_AclEntry_MatchedPacketsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_MatchedPacketsPath extracts the value of the leaf MatchedPackets from its parent oc.Acl_AclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertAcl_AclSet_AclEntry_MatchedPacketsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedPackets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
