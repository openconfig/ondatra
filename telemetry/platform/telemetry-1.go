package platform

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

// Lookup fetches the value at /openconfig-platform/components/component/fan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_FanPath) Lookup(t testing.TB) *oc.QualifiedComponent_Fan {
	t.Helper()
	goStruct := &oc.Component_Fan{}
	md, ok := oc.Lookup(t, n, "Component_Fan", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Fan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/fan with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_FanPath) Get(t testing.TB) *oc.Component_Fan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/fan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_FanPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Fan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Fan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Fan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Fan", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Fan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/fan with a ONCE subscription.
func (n *Component_FanPathAny) Get(t testing.TB) []*oc.Component_Fan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Fan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/fan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_FanPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Fan {
	t.Helper()
	c := &oc.CollectionComponent_Fan{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Fan) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Fan{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Fan)))
		return false
	})
	return c
}

func watch_Component_FanPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Fan) bool) *oc.Component_FanWatcher {
	t.Helper()
	w := &oc.Component_FanWatcher{}
	gs := &oc.Component_Fan{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Fan", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Fan{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Fan)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/fan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_FanPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Fan) bool) *oc.Component_FanWatcher {
	t.Helper()
	return watch_Component_FanPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/fan with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_FanPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Fan) *oc.QualifiedComponent_Fan {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Fan) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/fan failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/fan to the batch object.
func (n *Component_FanPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/fan with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_FanPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Fan {
	t.Helper()
	c := &oc.CollectionComponent_Fan{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Fan) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_FanPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Fan) bool) *oc.Component_FanWatcher {
	t.Helper()
	w := &oc.Component_FanWatcher{}
	structs := map[string]*oc.Component_Fan{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Fan{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Fan", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Fan{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Fan)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/fan with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_FanPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Fan) bool) *oc.Component_FanWatcher {
	t.Helper()
	return watch_Component_FanPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/fan to the batch object.
func (n *Component_FanPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/fan/state/speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Fan_SpeedPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Component_Fan{}
	md, ok := oc.Lookup(t, n, "Component_Fan", goStruct, true, false)
	if ok {
		return convertComponent_Fan_SpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/fan/state/speed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Fan_SpeedPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/fan/state/speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Fan_SpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Fan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Fan", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Fan_SpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/fan/state/speed with a ONCE subscription.
func (n *Component_Fan_SpeedPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/fan/state/speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Fan_SpeedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Fan_SpeedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Component_Fan{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Fan", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Fan_SpeedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/fan/state/speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Fan_SpeedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_Fan_SpeedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/fan/state/speed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Fan_SpeedPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/fan/state/speed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/fan/state/speed to the batch object.
func (n *Component_Fan_SpeedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/fan/state/speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Fan_SpeedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Fan_SpeedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Component_Fan{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Fan{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Fan", structs[pre], queryPath, true, false)
			qv := convertComponent_Fan_SpeedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/fan/state/speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Fan_SpeedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_Fan_SpeedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/fan/state/speed to the batch object.
func (n *Component_Fan_SpeedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Fan_SpeedPath extracts the value of the leaf Speed from its parent oc.Component_Fan
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertComponent_Fan_SpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Fan) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Speed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/firmware-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_FirmwareVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_FirmwareVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/firmware-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_FirmwareVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/firmware-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_FirmwareVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_FirmwareVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/firmware-version with a ONCE subscription.
func (n *Component_FirmwareVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/firmware-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_FirmwareVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_FirmwareVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_FirmwareVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/firmware-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_FirmwareVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_FirmwareVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/firmware-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_FirmwareVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/firmware-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/firmware-version to the batch object.
func (n *Component_FirmwareVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/firmware-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_FirmwareVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_FirmwareVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_FirmwareVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/firmware-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_FirmwareVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_FirmwareVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/firmware-version to the batch object.
func (n *Component_FirmwareVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_FirmwareVersionPath extracts the value of the leaf FirmwareVersion from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_FirmwareVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.FirmwareVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/hardware-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_HardwareVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_HardwareVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/hardware-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_HardwareVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/hardware-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_HardwareVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_HardwareVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/hardware-version with a ONCE subscription.
func (n *Component_HardwareVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/hardware-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_HardwareVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_HardwareVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_HardwareVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/hardware-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_HardwareVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_HardwareVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/hardware-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_HardwareVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/hardware-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/hardware-version to the batch object.
func (n *Component_HardwareVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/hardware-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_HardwareVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_HardwareVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_HardwareVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/hardware-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_HardwareVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_HardwareVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/hardware-version to the batch object.
func (n *Component_HardwareVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_HardwareVersionPath extracts the value of the leaf HardwareVersion from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_HardwareVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.HardwareVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/id with a ONCE subscription.
func (n *Component_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/id to the batch object.
func (n *Component_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/id to the batch object.
func (n *Component_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IdPath extracts the value of the leaf Id from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuitPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuitPath) Get(t testing.TB) *oc.Component_IntegratedCircuit {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuitPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit with a ONCE subscription.
func (n *Component_IntegratedCircuitPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit) bool) *oc.Component_IntegratedCircuitWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuitWatcher{}
	gs := &oc.Component_IntegratedCircuit{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit) bool) *oc.Component_IntegratedCircuitWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuitPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit) *oc.QualifiedComponent_IntegratedCircuit {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit to the batch object.
func (n *Component_IntegratedCircuitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuitPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit) bool) *oc.Component_IntegratedCircuitWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuitWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit) bool) *oc.Component_IntegratedCircuitWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuitPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit to the batch object.
func (n *Component_IntegratedCircuitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a ONCE subscription.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_BackplaneFacingCapacity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit_BackplaneFacingCapacity)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool) *oc.Component_IntegratedCircuit_BackplaneFacingCapacityWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_BackplaneFacingCapacityWatcher{}
	gs := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool) *oc.Component_IntegratedCircuit_BackplaneFacingCapacityWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool) *oc.Component_IntegratedCircuit_BackplaneFacingCapacityWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_BackplaneFacingCapacityWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) bool) *oc.Component_IntegratedCircuit_BackplaneFacingCapacityWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a ONCE subscription.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/available-pct to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath extracts the value of the leaf AvailablePct from its parent oc.Component_IntegratedCircuit_BackplaneFacingCapacity
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertComponent_IntegratedCircuit_BackplaneFacingCapacity_AvailablePctPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AvailablePct
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a ONCE subscription.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/consumed-capacity to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath extracts the value of the leaf ConsumedCapacity from its parent oc.Component_IntegratedCircuit_BackplaneFacingCapacity
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_BackplaneFacingCapacity_ConsumedCapacityPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConsumedCapacity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a ONCE subscription.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total-operational-capacity to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath extracts the value of the leaf TotalOperationalCapacity from its parent oc.Component_IntegratedCircuit_BackplaneFacingCapacity
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalOperationalCapacityPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TotalOperationalCapacity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a ONCE subscription.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_BackplaneFacingCapacity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_BackplaneFacingCapacity", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/backplane-facing-capacity/state/total to the batch object.
func (n *Component_IntegratedCircuit_BackplaneFacingCapacity_TotalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalPath extracts the value of the leaf Total from its parent oc.Component_IntegratedCircuit_BackplaneFacingCapacity
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_BackplaneFacingCapacity_TotalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_BackplaneFacingCapacity) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Total
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_MemoryPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Memory {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Memory{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Memory", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Memory{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_MemoryPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Memory {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_MemoryPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Memory {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Memory
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Memory", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Memory{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/memory with a ONCE subscription.
func (n *Component_IntegratedCircuit_MemoryPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Memory {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Memory
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_MemoryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Memory {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Memory{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Memory) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit_Memory{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit_Memory)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_MemoryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Memory) bool) *oc.Component_IntegratedCircuit_MemoryWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_MemoryWatcher{}
	gs := &oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Memory", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Memory{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Memory)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_MemoryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Memory) bool) *oc.Component_IntegratedCircuit_MemoryWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_MemoryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/memory with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_MemoryPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit_Memory) *oc.QualifiedComponent_IntegratedCircuit_Memory {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit_Memory) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/memory failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory to the batch object.
func (n *Component_IntegratedCircuit_MemoryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_MemoryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Memory {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Memory{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Memory) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_MemoryPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Memory) bool) *oc.Component_IntegratedCircuit_MemoryWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_MemoryWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Memory", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit_Memory{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Memory)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_MemoryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Memory) bool) *oc.Component_IntegratedCircuit_MemoryWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_MemoryPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory to the batch object.
func (n *Component_IntegratedCircuit_MemoryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Memory{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Memory", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a ONCE subscription.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Memory", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Memory", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/corrected-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_CorrectedParityErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Memory_CorrectedParityErrorsPath extracts the value of the leaf CorrectedParityErrors from its parent oc.Component_IntegratedCircuit_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Memory_CorrectedParityErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CorrectedParityErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Memory{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Memory", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Memory_TotalParityErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Memory_TotalParityErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a ONCE subscription.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_TotalParityErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Memory", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Memory_TotalParityErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_TotalParityErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Memory", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Memory_TotalParityErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/total-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_TotalParityErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Memory_TotalParityErrorsPath extracts the value of the leaf TotalParityErrors from its parent oc.Component_IntegratedCircuit_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Memory_TotalParityErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TotalParityErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Memory{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Memory", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Memory", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Memory", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath extracts the value of the leaf UncorrectedParityErrors from its parent oc.Component_IntegratedCircuit_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UncorrectedParityErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_NodeIdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_NodeIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_NodeIdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_NodeIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_NodeIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_NodeIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_NodeIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_NodeIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_NodeIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_NodeIdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/state/node-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/state/node-id to the batch object.
func (n *Component_IntegratedCircuit_NodeIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_NodeIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_NodeIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_NodeIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/state/node-id to the batch object.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_NodeIdPath extracts the value of the leaf NodeId from its parent oc.Component_IntegratedCircuit
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_NodeIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.NodeId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_UtilizationPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_UtilizationPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Utilization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Utilization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Utilization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Utilization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_UtilizationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit_Utilization)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_UtilizationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_UtilizationWatcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_UtilizationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_UtilizationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_UtilizationPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit_Utilization) *oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization to the batch object.
func (n *Component_IntegratedCircuit_UtilizationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_UtilizationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_UtilizationWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_UtilizationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization to the batch object.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Utilization_Resource
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization_Resource{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit_Utilization_Resource)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_ResourcePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_Utilization_ResourceWatcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_ResourcePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource to the batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization_Resource{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_ResourcePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_Utilization_ResourceWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_ResourcePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource to the batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath extracts the value of the leaf Committed from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Committed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_FreePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_FreePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_FreePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_FreePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_FreePath extracts the value of the leaf Free from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Free
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath extracts the value of the leaf HighWatermark from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.HighWatermark
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath extracts the value of the leaf LastHighWatermark from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastHighWatermark
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath extracts the value of the leaf MaxLimit from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_NamePath extracts the value of the leaf Name from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_UsedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_UsedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_UsedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_UsedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath extracts the value of the leaf Used from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Used
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastRebootReasonPath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_LastRebootReasonPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastRebootReasonPath) Get(t testing.TB) oc.E_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastRebootReasonPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastRebootReasonPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription.
func (n *Component_LastRebootReasonPathAny) Get(t testing.TB) []oc.E_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_COMPONENT_REBOOT_REASON
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootReasonPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootReasonPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastRebootReasonPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootReasonPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	return watch_Component_LastRebootReasonPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastRebootReasonPath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_COMPONENT_REBOOT_REASON) *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-reboot-reason failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-reason to the batch object.
func (n *Component_LastRebootReasonPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootReasonPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootReasonPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_LastRebootReasonPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootReasonPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	return watch_Component_LastRebootReasonPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-reason to the batch object.
func (n *Component_LastRebootReasonPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastRebootReasonPath extracts the value of the leaf LastRebootReason from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON.
func convertComponent_LastRebootReasonPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON{
		Metadata: md,
	}
	val := parent.LastRebootReason
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastRebootTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_LastRebootTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastRebootTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastRebootTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastRebootTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription.
func (n *Component_LastRebootTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastRebootTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_LastRebootTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastRebootTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-reboot-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-time to the batch object.
func (n *Component_LastRebootTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_LastRebootTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_LastRebootTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-time to the batch object.
func (n *Component_LastRebootTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastRebootTimePath extracts the value of the leaf LastRebootTime from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_LastRebootTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastRebootTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastSwitchoverReasonPath) Lookup(t testing.TB) *oc.QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	goStruct := &oc.Component_LastSwitchoverReason{}
	md, ok := oc.Lookup(t, n, "Component_LastSwitchoverReason", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastSwitchoverReasonPath) Get(t testing.TB) *oc.Component_LastSwitchoverReason {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastSwitchoverReasonPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_LastSwitchoverReason
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_LastSwitchoverReason{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_LastSwitchoverReason", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription.
func (n *Component_LastSwitchoverReasonPathAny) Get(t testing.TB) []*oc.Component_LastSwitchoverReason {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_LastSwitchoverReason
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReasonPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_LastSwitchoverReason {
	t.Helper()
	c := &oc.CollectionComponent_LastSwitchoverReason{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_LastSwitchoverReason) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_LastSwitchoverReason)))
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReasonPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	w := &oc.Component_LastSwitchoverReasonWatcher{}
	gs := &oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_LastSwitchoverReason", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_LastSwitchoverReason)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReasonPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReasonPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastSwitchoverReasonPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_LastSwitchoverReason) *oc.QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_LastSwitchoverReason) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-switchover-reason failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason to the batch object.
func (n *Component_LastSwitchoverReasonPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReasonPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_LastSwitchoverReason {
	t.Helper()
	c := &oc.CollectionComponent_LastSwitchoverReason{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_LastSwitchoverReason) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReasonPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	w := &oc.Component_LastSwitchoverReasonWatcher{}
	structs := map[string]*oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_LastSwitchoverReason{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_LastSwitchoverReason", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_LastSwitchoverReason{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_LastSwitchoverReason)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReasonPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReasonPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason to the batch object.
func (n *Component_LastSwitchoverReasonPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastSwitchoverReason_DetailsPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_LastSwitchoverReason{}
	md, ok := oc.Lookup(t, n, "Component_LastSwitchoverReason", goStruct, true, false)
	if ok {
		return convertComponent_LastSwitchoverReason_DetailsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastSwitchoverReason_DetailsPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_LastSwitchoverReason{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_LastSwitchoverReason", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastSwitchoverReason_DetailsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReason_DetailsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReason_DetailsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_LastSwitchoverReason", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastSwitchoverReason_DetailsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReason_DetailsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReason_DetailsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastSwitchoverReason_DetailsPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-switchover-reason/details failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason/details to the batch object.
func (n *Component_LastSwitchoverReason_DetailsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReason_DetailsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_LastSwitchoverReason{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_LastSwitchoverReason", structs[pre], queryPath, true, false)
			qv := convertComponent_LastSwitchoverReason_DetailsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReason_DetailsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason/details to the batch object.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastSwitchoverReason_DetailsPath extracts the value of the leaf Details from its parent oc.Component_LastSwitchoverReason
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_LastSwitchoverReason_DetailsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_LastSwitchoverReason) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Details
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
