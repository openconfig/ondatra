package opentrafficgeneratorport

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/otgtelemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Port{}
	md, ok := oc.Lookup(t, n, "Port", goStruct, true, false)
	if ok {
		return convertPort_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertPort_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/name with a ONCE subscription.
func (n *Port_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Port_NamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/name to the batch object.
func (n *Port_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Port{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port", structs[pre], queryPath, true, false)
			qv := convertPort_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Port_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/name to the batch object.
func (n *Port_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_NamePath extracts the value of the leaf Name from its parent oc.Port
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertPort_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Port) *oc.QualifiedString {
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

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/out-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_OutRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Port{}
	md, ok := oc.Lookup(t, n, "Port", goStruct, true, false)
	if ok {
		return convertPort_OutRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/out-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_OutRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/out-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_OutRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertPort_OutRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/out-rate with a ONCE subscription.
func (n *Port_OutRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_OutRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_OutRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_OutRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_OutRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Port_OutRatePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/out-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_OutRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/out-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/out-rate to the batch object.
func (n *Port_OutRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_OutRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_OutRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Port{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port", structs[pre], queryPath, true, false)
			qv := convertPort_OutRatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_OutRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Port_OutRatePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/out-rate to the batch object.
func (n *Port_OutRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_OutRatePath extracts the value of the leaf OutRate from its parent oc.Port
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertPort_OutRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Port) *oc.QualifiedFloat32 {
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
