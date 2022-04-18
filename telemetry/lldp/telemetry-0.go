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

// Lookup fetches the value at /openconfig-lldp/lldp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpPath) Lookup(t testing.TB) *oc.QualifiedLldp {
	t.Helper()
	goStruct := &oc.Lldp{}
	md, ok := oc.Lookup(t, n, "Lldp", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpPath) Get(t testing.TB) *oc.Lldp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpPathAny) Lookup(t testing.TB) []*oc.QualifiedLldp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp with a ONCE subscription.
func (n *LldpPathAny) Get(t testing.TB) []*oc.Lldp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp {
	t.Helper()
	c := &oc.CollectionLldp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp)))
		return false
	})
	return c
}

func watch_LldpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp) bool) *oc.LldpWatcher {
	t.Helper()
	w := &oc.LldpWatcher{}
	gs := &oc.Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldp{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp) bool) *oc.LldpWatcher {
	t.Helper()
	return watch_LldpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpPath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp) *oc.QualifiedLldp {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp to the batch object.
func (n *LldpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp {
	t.Helper()
	c := &oc.CollectionLldp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp) bool) *oc.LldpWatcher {
	t.Helper()
	w := &oc.LldpWatcher{}
	structs := map[string]*oc.Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldp{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp) bool) *oc.LldpWatcher {
	t.Helper()
	return watch_LldpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp to the batch object.
func (n *LldpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lldp/lldp/state/chassis-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_ChassisIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp{}
	md, ok := oc.Lookup(t, n, "Lldp", goStruct, true, false)
	if ok {
		return convertLldp_ChassisIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/state/chassis-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_ChassisIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/state/chassis-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_ChassisIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_ChassisIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/state/chassis-id with a ONCE subscription.
func (n *Lldp_ChassisIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_ChassisIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_ChassisIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_ChassisIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_ChassisIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_ChassisIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/state/chassis-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_ChassisIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/state/chassis-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/state/chassis-id to the batch object.
func (n *Lldp_ChassisIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_ChassisIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_ChassisIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp", structs[pre], queryPath, true, false)
			qv := convertLldp_ChassisIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_ChassisIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_ChassisIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/state/chassis-id to the batch object.
func (n *Lldp_ChassisIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_ChassisIdPath extracts the value of the leaf ChassisId from its parent oc.Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_ChassisIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ChassisId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/state/chassis-id-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_ChassisIdTypePath) Lookup(t testing.TB) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	goStruct := &oc.Lldp{}
	md, ok := oc.Lookup(t, n, "Lldp", goStruct, true, false)
	if ok {
		return convertLldp_ChassisIdTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/state/chassis-id-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_ChassisIdTypePath) Get(t testing.TB) oc.E_LldpTypes_ChassisIdType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/state/chassis-id-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_ChassisIdTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpTypes_ChassisIdType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_ChassisIdTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/state/chassis-id-type with a ONCE subscription.
func (n *Lldp_ChassisIdTypePathAny) Get(t testing.TB) []oc.E_LldpTypes_ChassisIdType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpTypes_ChassisIdType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_ChassisIdTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_ChassisIdTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_ChassisIdTypeWatcher{}
	gs := &oc.Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_ChassisIdTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_ChassisIdTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	return watch_Lldp_ChassisIdTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/state/chassis-id-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_ChassisIdTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpTypes_ChassisIdType) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/state/chassis-id-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/state/chassis-id-type to the batch object.
func (n *Lldp_ChassisIdTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_ChassisIdTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_ChassisIdTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_ChassisIdTypeWatcher{}
	structs := map[string]*oc.Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp", structs[pre], queryPath, true, false)
			qv := convertLldp_ChassisIdTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_ChassisIdTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	return watch_Lldp_ChassisIdTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/state/chassis-id-type to the batch object.
func (n *Lldp_ChassisIdTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_ChassisIdTypePath extracts the value of the leaf ChassisIdType from its parent oc.Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpTypes_ChassisIdType.
func convertLldp_ChassisIdTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	qv := &oc.QualifiedE_LldpTypes_ChassisIdType{
		Metadata: md,
	}
	val := parent.ChassisIdType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_CountersPath) Lookup(t testing.TB) *oc.QualifiedLldp_Counters {
	t.Helper()
	goStruct := &oc.Lldp_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldp_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_CountersPath) Get(t testing.TB) *oc.Lldp_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/state/counters with a ONCE subscription.
func (n *Lldp_CountersPathAny) Get(t testing.TB) []*oc.Lldp_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Counters {
	t.Helper()
	c := &oc.CollectionLldp_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldp_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lldp_Counters)))
		return false
	})
	return c
}

func watch_Lldp_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Counters) bool) *oc.Lldp_CountersWatcher {
	t.Helper()
	w := &oc.Lldp_CountersWatcher{}
	gs := &oc.Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldp_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Counters) bool) *oc.Lldp_CountersWatcher {
	t.Helper()
	return watch_Lldp_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Lldp_Counters) *oc.QualifiedLldp_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldp_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/state/counters to the batch object.
func (n *Lldp_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldp_Counters {
	t.Helper()
	c := &oc.CollectionLldp_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldp_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldp_Counters) bool) *oc.Lldp_CountersWatcher {
	t.Helper()
	w := &oc.Lldp_CountersWatcher{}
	structs := map[string]*oc.Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldp_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldp_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldp_Counters) bool) *oc.Lldp_CountersWatcher {
	t.Helper()
	return watch_Lldp_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/state/counters to the batch object.
func (n *Lldp_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lldp/lldp/state/counters/entries-aged-out with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Counters_EntriesAgedOutPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Counters_EntriesAgedOutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/state/counters/entries-aged-out with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Counters_EntriesAgedOutPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Counters_EntriesAgedOutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Counters_EntriesAgedOutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a ONCE subscription.
func (n *Lldp_Counters_EntriesAgedOutPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Counters_EntriesAgedOutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Counters_EntriesAgedOutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Counters_EntriesAgedOutPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Counters_EntriesAgedOutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Counters_EntriesAgedOutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Counters_EntriesAgedOutPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/state/counters/entries-aged-out failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/state/counters/entries-aged-out to the batch object.
func (n *Lldp_Counters_EntriesAgedOutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Counters_EntriesAgedOutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Counters_EntriesAgedOutPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Counters_EntriesAgedOutPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/counters/entries-aged-out with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Counters_EntriesAgedOutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Counters_EntriesAgedOutPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/state/counters/entries-aged-out to the batch object.
func (n *Lldp_Counters_EntriesAgedOutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Counters_EntriesAgedOutPath extracts the value of the leaf EntriesAgedOut from its parent oc.Lldp_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Counters_EntriesAgedOutPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.EntriesAgedOut
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/state/counters/frame-discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Counters_FrameDiscardPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp_Counters{}
	md, ok := oc.Lookup(t, n, "Lldp_Counters", goStruct, true, false)
	if ok {
		return convertLldp_Counters_FrameDiscardPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/state/counters/frame-discard with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Counters_FrameDiscardPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/state/counters/frame-discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Counters_FrameDiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Counters_FrameDiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/state/counters/frame-discard with a ONCE subscription.
func (n *Lldp_Counters_FrameDiscardPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/counters/frame-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Counters_FrameDiscardPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Counters_FrameDiscardPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Counters_FrameDiscardPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/counters/frame-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Counters_FrameDiscardPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Counters_FrameDiscardPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/state/counters/frame-discard with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Counters_FrameDiscardPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/state/counters/frame-discard failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/state/counters/frame-discard to the batch object.
func (n *Lldp_Counters_FrameDiscardPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/state/counters/frame-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Counters_FrameDiscardPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Counters_FrameDiscardPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Counters", structs[pre], queryPath, true, false)
			qv := convertLldp_Counters_FrameDiscardPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/state/counters/frame-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Counters_FrameDiscardPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lldp_Counters_FrameDiscardPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/state/counters/frame-discard to the batch object.
func (n *Lldp_Counters_FrameDiscardPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Counters_FrameDiscardPath extracts the value of the leaf FrameDiscard from its parent oc.Lldp_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_Counters_FrameDiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FrameDiscard
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
