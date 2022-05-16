package interfaces

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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/hardware-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HardwarePortPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_HardwarePortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/hardware-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HardwarePortPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HardwarePortPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_HardwarePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a ONCE subscription.
func (n *Interface_HardwarePortPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HardwarePortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HardwarePortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_HardwarePortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HardwarePortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_HardwarePortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_HardwarePortPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/hardware-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/hardware-port to the batch object.
func (n *Interface_HardwarePortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HardwarePortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HardwarePortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_HardwarePortPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/hardware-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HardwarePortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_HardwarePortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/hardware-port to the batch object.
func (n *Interface_HardwarePortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_HardwarePortPath extracts the value of the leaf HardwarePort from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_HardwarePortPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.HardwarePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HoldTimePath) Lookup(t testing.TB) *oc.QualifiedInterface_HoldTime {
	t.Helper()
	goStruct := &oc.Interface_HoldTime{}
	md, ok := oc.Lookup(t, n, "Interface_HoldTime", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_HoldTime{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HoldTimePath) Get(t testing.TB) *oc.Interface_HoldTime {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HoldTimePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_HoldTime {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_HoldTime
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_HoldTime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_HoldTime", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_HoldTime{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/hold-time with a ONCE subscription.
func (n *Interface_HoldTimePathAny) Get(t testing.TB) []*oc.Interface_HoldTime {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_HoldTime
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/hold-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HoldTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_HoldTime {
	t.Helper()
	c := &oc.CollectionInterface_HoldTime{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_HoldTime) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_HoldTime{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_HoldTime)))
		return false
	})
	return c
}

func watch_Interface_HoldTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_HoldTime) bool) *oc.Interface_HoldTimeWatcher {
	t.Helper()
	w := &oc.Interface_HoldTimeWatcher{}
	gs := &oc.Interface_HoldTime{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_HoldTime", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_HoldTime{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_HoldTime)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/hold-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HoldTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_HoldTime) bool) *oc.Interface_HoldTimeWatcher {
	t.Helper()
	return watch_Interface_HoldTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/hold-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_HoldTimePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_HoldTime) *oc.QualifiedInterface_HoldTime {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_HoldTime) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/hold-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/hold-time to the batch object.
func (n *Interface_HoldTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/hold-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HoldTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_HoldTime {
	t.Helper()
	c := &oc.CollectionInterface_HoldTime{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_HoldTime) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HoldTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_HoldTime) bool) *oc.Interface_HoldTimeWatcher {
	t.Helper()
	w := &oc.Interface_HoldTimeWatcher{}
	structs := map[string]*oc.Interface_HoldTime{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_HoldTime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_HoldTime", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_HoldTime{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_HoldTime)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/hold-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HoldTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_HoldTime) bool) *oc.Interface_HoldTimeWatcher {
	t.Helper()
	return watch_Interface_HoldTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/hold-time to the batch object.
func (n *Interface_HoldTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HoldTime_DownPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_HoldTime{}
	md, ok := oc.Lookup(t, n, "Interface_HoldTime", goStruct, true, false)
	if ok {
		return convertInterface_HoldTime_DownPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetDown())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HoldTime_DownPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HoldTime_DownPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_HoldTime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_HoldTime", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_HoldTime_DownPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a ONCE subscription.
func (n *Interface_HoldTime_DownPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HoldTime_DownPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HoldTime_DownPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_HoldTime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_HoldTime", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_HoldTime_DownPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HoldTime_DownPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_HoldTime_DownPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_HoldTime_DownPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/hold-time/state/down failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/hold-time/state/down to the batch object.
func (n *Interface_HoldTime_DownPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HoldTime_DownPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HoldTime_DownPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_HoldTime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_HoldTime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_HoldTime", structs[pre], queryPath, true, false)
			qv := convertInterface_HoldTime_DownPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/down with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HoldTime_DownPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_HoldTime_DownPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/hold-time/state/down to the batch object.
func (n *Interface_HoldTime_DownPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_HoldTime_DownPath extracts the value of the leaf Down from its parent oc.Interface_HoldTime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_HoldTime_DownPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_HoldTime) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Down
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_HoldTime_UpPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_HoldTime{}
	md, ok := oc.Lookup(t, n, "Interface_HoldTime", goStruct, true, false)
	if ok {
		return convertInterface_HoldTime_UpPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetUp())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_HoldTime_UpPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_HoldTime_UpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_HoldTime{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_HoldTime", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_HoldTime_UpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a ONCE subscription.
func (n *Interface_HoldTime_UpPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HoldTime_UpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HoldTime_UpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_HoldTime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_HoldTime", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_HoldTime_UpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HoldTime_UpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_HoldTime_UpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_HoldTime_UpPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/hold-time/state/up failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/hold-time/state/up to the batch object.
func (n *Interface_HoldTime_UpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_HoldTime_UpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_HoldTime_UpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_HoldTime{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_HoldTime{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_HoldTime", structs[pre], queryPath, true, false)
			qv := convertInterface_HoldTime_UpPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/hold-time/state/up with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_HoldTime_UpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_HoldTime_UpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/hold-time/state/up to the batch object.
func (n *Interface_HoldTime_UpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_HoldTime_UpPath extracts the value of the leaf Up from its parent oc.Interface_HoldTime
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_HoldTime_UpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_HoldTime) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Up
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_IdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_IdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/id with a ONCE subscription.
func (n *Interface_IdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_IdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/id to the batch object.
func (n *Interface_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_IdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/id to the batch object.
func (n *Interface_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_IdPath extracts the value of the leaf Id from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/ifindex with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_IfindexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_IfindexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/ifindex with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_IfindexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/ifindex with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_IfindexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_IfindexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/ifindex with a ONCE subscription.
func (n *Interface_IfindexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/ifindex with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_IfindexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_IfindexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_IfindexPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/ifindex with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_IfindexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_IfindexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/ifindex with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_IfindexPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/ifindex failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/ifindex to the batch object.
func (n *Interface_IfindexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/ifindex with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_IfindexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_IfindexPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_IfindexPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/ifindex with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_IfindexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_IfindexPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/ifindex to the batch object.
func (n *Interface_IfindexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_IfindexPath extracts the value of the leaf Ifindex from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_IfindexPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Ifindex
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/in-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_InRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_InRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/in-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_InRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/in-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_InRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_InRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/in-rate with a ONCE subscription.
func (n *Interface_InRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/in-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_InRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_InRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_InRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/in-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_InRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Interface_InRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/in-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_InRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/in-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/in-rate to the batch object.
func (n *Interface_InRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/in-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_InRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_InRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_InRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/in-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_InRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Interface_InRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/in-rate to the batch object.
func (n *Interface_InRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_InRatePath extracts the value of the leaf InRate from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertInterface_InRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.InRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/last-change with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_LastChangePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_LastChangePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/last-change with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_LastChangePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/last-change with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_LastChangePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_LastChangePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/last-change with a ONCE subscription.
func (n *Interface_LastChangePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/last-change with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_LastChangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_LastChangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_LastChangePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/last-change with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_LastChangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_LastChangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/last-change with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_LastChangePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/last-change failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/last-change to the batch object.
func (n *Interface_LastChangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/last-change with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_LastChangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_LastChangePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_LastChangePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/last-change with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_LastChangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_LastChangePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/last-change to the batch object.
func (n *Interface_LastChangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_LastChangePath extracts the value of the leaf LastChange from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_LastChangePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastChange
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/logical with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_LogicalPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_LogicalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/logical with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_LogicalPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/logical with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_LogicalPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_LogicalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/logical with a ONCE subscription.
func (n *Interface_LogicalPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/logical with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_LogicalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_LogicalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_LogicalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/logical with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_LogicalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_LogicalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/logical with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_LogicalPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/logical failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/logical to the batch object.
func (n *Interface_LogicalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/logical with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_LogicalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_LogicalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_LogicalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/logical with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_LogicalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_LogicalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/logical to the batch object.
func (n *Interface_LogicalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_LogicalPath extracts the value of the leaf Logical from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_LogicalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Logical
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_LoopbackModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_LoopbackModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetLoopbackMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_LoopbackModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_LoopbackModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_LoopbackModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a ONCE subscription.
func (n *Interface_LoopbackModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_LoopbackModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_LoopbackModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_LoopbackModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_LoopbackModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_LoopbackModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_LoopbackModePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/loopback-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/loopback-mode to the batch object.
func (n *Interface_LoopbackModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_LoopbackModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_LoopbackModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_LoopbackModePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/loopback-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_LoopbackModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_LoopbackModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/loopback-mode to the batch object.
func (n *Interface_LoopbackModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_LoopbackModePath extracts the value of the leaf LoopbackMode from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_LoopbackModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.LoopbackMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/management with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_ManagementPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_ManagementPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/management with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_ManagementPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/management with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_ManagementPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_ManagementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/management with a ONCE subscription.
func (n *Interface_ManagementPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/management with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_ManagementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_ManagementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_ManagementPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/management with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_ManagementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_ManagementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/management with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_ManagementPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/management failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/management to the batch object.
func (n *Interface_ManagementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/management with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_ManagementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_ManagementPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_ManagementPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/management with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_ManagementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_ManagementPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/management to the batch object.
func (n *Interface_ManagementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_ManagementPath extracts the value of the leaf Management from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_ManagementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Management
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/mtu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_MtuPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/mtu with a ONCE subscription.
func (n *Interface_MtuPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_MtuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_MtuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_MtuPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_MtuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_MtuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/mtu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_MtuPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/mtu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/mtu to the batch object.
func (n *Interface_MtuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_MtuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_MtuPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_MtuPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_MtuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_MtuPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/mtu to the batch object.
func (n *Interface_MtuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Mtu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/name with a ONCE subscription.
func (n *Interface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/name to the batch object.
func (n *Interface_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/name to the batch object.
func (n *Interface_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_NamePath extracts the value of the leaf Name from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/oper-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_OperStatusPath) Lookup(t testing.TB) *oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_OperStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/oper-status with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_OperStatusPath) Get(t testing.TB) oc.E_Interface_OperStatus {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/oper-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_OperStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Interface_OperStatus
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_OperStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/oper-status with a ONCE subscription.
func (n *Interface_OperStatusPathAny) Get(t testing.TB) []oc.E_Interface_OperStatus {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Interface_OperStatus
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/oper-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_OperStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Interface_OperStatus {
	t.Helper()
	c := &oc.CollectionE_Interface_OperStatus{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Interface_OperStatus) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_OperStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	w := &oc.E_Interface_OperStatusWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_OperStatusPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Interface_OperStatus)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/oper-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_OperStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	return watch_Interface_OperStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/oper-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_OperStatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_Interface_OperStatus) *oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Interface_OperStatus) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/oper-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/oper-status to the batch object.
func (n *Interface_OperStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/oper-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_OperStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Interface_OperStatus {
	t.Helper()
	c := &oc.CollectionE_Interface_OperStatus{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Interface_OperStatus) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_OperStatusPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	w := &oc.E_Interface_OperStatusWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_OperStatusPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Interface_OperStatus)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/oper-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_OperStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Interface_OperStatus) bool) *oc.E_Interface_OperStatusWatcher {
	t.Helper()
	return watch_Interface_OperStatusPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/oper-status to the batch object.
func (n *Interface_OperStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_OperStatusPath extracts the value of the leaf OperStatus from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Interface_OperStatus.
func convertInterface_OperStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedE_Interface_OperStatus {
	t.Helper()
	qv := &oc.QualifiedE_Interface_OperStatus{
		Metadata: md,
	}
	val := parent.OperStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/out-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_OutRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_OutRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/out-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_OutRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/out-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_OutRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_OutRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/out-rate with a ONCE subscription.
func (n *Interface_OutRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_OutRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_OutRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_OutRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_OutRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Interface_OutRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/out-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_OutRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/out-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/out-rate to the batch object.
func (n *Interface_OutRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_OutRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_OutRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_OutRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_OutRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Interface_OutRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/out-rate to the batch object.
func (n *Interface_OutRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_OutRatePath extracts the value of the leaf OutRate from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertInterface_OutRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.OutRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/physical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_PhysicalChannelPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_PhysicalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/physical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_PhysicalChannelPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_PhysicalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_PhysicalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a ONCE subscription.
func (n *Interface_PhysicalChannelPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_PhysicalChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_PhysicalChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_PhysicalChannelPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_PhysicalChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_PhysicalChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_PhysicalChannelPath) Await(t testing.TB, timeout time.Duration, val []uint16) *oc.QualifiedUint16Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/physical-channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/physical-channel to the batch object.
func (n *Interface_PhysicalChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_PhysicalChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_PhysicalChannelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, true, false)
			qv := convertInterface_PhysicalChannelPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/physical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_PhysicalChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_PhysicalChannelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/physical-channel to the batch object.
func (n *Interface_PhysicalChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_PhysicalChannelPath extracts the value of the leaf PhysicalChannel from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_PhysicalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.PhysicalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlanPath) Get(t testing.TB) *oc.Interface_RoutedVlan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a ONCE subscription.
func (n *Interface_RoutedVlanPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlanPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlanPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan) bool) *oc.Interface_RoutedVlanWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlanWatcher{}
	gs := &oc.Interface_RoutedVlan{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlanPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan) bool) *oc.Interface_RoutedVlanWatcher {
	t.Helper()
	return watch_Interface_RoutedVlanPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlanPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan) *oc.QualifiedInterface_RoutedVlan {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan to the batch object.
func (n *Interface_RoutedVlanPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlanPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlanPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan) bool) *oc.Interface_RoutedVlanWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlanWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlanPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan) bool) *oc.Interface_RoutedVlanWatcher {
	t.Helper()
	return watch_Interface_RoutedVlanPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan to the batch object.
func (n *Interface_RoutedVlanPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4Path) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4PathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4 {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4) bool) *oc.Interface_RoutedVlan_Ipv4Watcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4) bool) *oc.Interface_RoutedVlan_Ipv4Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4Path) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedInterface_RoutedVlan_Ipv4 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 to the batch object.
func (n *Interface_RoutedVlan_Ipv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4 {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4) bool) *oc.Interface_RoutedVlan_Ipv4Watcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4) bool) *oc.Interface_RoutedVlan_Ipv4Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4 to the batch object.
func (n *Interface_RoutedVlan_Ipv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Address)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool) *oc.Interface_RoutedVlan_Ipv4_AddressWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_AddressWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool) *oc.Interface_RoutedVlan_Ipv4_AddressWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Address) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address to the batch object.
func (n *Interface_RoutedVlan_Ipv4_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool) *oc.Interface_RoutedVlan_Ipv4_AddressWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_AddressWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address) bool) *oc.Interface_RoutedVlan_Ipv4_AddressWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address to the batch object.
func (n *Interface_RoutedVlan_Ipv4_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_IpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_IpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_IpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_IpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPath) Get(t testing.TB) oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_IpAddressOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_IpAddressOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_OriginPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_IpAddressOrigin) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_OriginPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_OriginPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_OriginPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_RoutedVlan_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_IpAddressOrigin.
func convertInterface_RoutedVlan_Ipv4_Address_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_IpAddressOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/state/prefix-length to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_RoutedVlan_Ipv4_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAcceptMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/accept-mode to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath extracts the value of the leaf AcceptMode from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AcceptModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AcceptMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAdvertisementInterval())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/advertisement-interval to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath extracts the value of the leaf AdvertisementInterval from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AdvertisementInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/current-priority to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath extracts the value of the leaf CurrentPriority from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_CurrentPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CurrentPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriorityDecrement())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath extracts the value of the leaf PriorityDecrement from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PriorityDecrement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath extracts the value of the leaf TrackInterface from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.TrackInterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPreemptDelay())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt-delay to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath extracts the value of the leaf PreemptDelay from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PreemptDelay
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPreempt())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/preempt to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath extracts the value of the leaf Preempt from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PreemptPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Preempt
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriority())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/priority to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath extracts the value of the leaf Priority from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-address to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath extracts the value of the leaf VirtualAddress from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.VirtualAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/addresses/address/vrrp/vrrp-group/state/virtual-router-id to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath extracts the value of the leaf VirtualRouterId from its parent oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv4_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.VirtualRouterId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_CountersPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_CountersPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_CountersPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Counters)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool) *oc.Interface_RoutedVlan_Ipv4_CountersWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_CountersWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool) *oc.Interface_RoutedVlan_Ipv4_CountersWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters to the batch object.
func (n *Interface_RoutedVlan_Ipv4_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool) *oc.Interface_RoutedVlan_Ipv4_CountersWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_CountersWatcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Counters) bool) *oc.Interface_RoutedVlan_Ipv4_CountersWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters to the batch object.
func (n *Interface_RoutedVlan_Ipv4_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_RoutedVlan_Ipv4_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath extracts the value of the leaf InDiscardedPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_InDiscardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InDiscardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
