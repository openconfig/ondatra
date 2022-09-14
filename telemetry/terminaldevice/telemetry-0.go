package terminaldevice

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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevicePath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice {
	t.Helper()
	goStruct := &oc.TerminalDevice{}
	md, ok := oc.Lookup(t, n, "TerminalDevice", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevicePath) Get(t testing.TB) *oc.TerminalDevice {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevicePathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device with a ONCE subscription.
func (n *TerminalDevicePathAny) Get(t testing.TB) []*oc.TerminalDevice {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevicePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice {
	t.Helper()
	c := &oc.CollectionTerminalDevice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice)))
		return false
	})
	return c
}

func watch_TerminalDevicePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice) bool) *oc.TerminalDeviceWatcher {
	t.Helper()
	w := &oc.TerminalDeviceWatcher{}
	gs := &oc.TerminalDevice{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevicePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice) bool) *oc.TerminalDeviceWatcher {
	t.Helper()
	return watch_TerminalDevicePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevicePath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice) *oc.QualifiedTerminalDevice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device to the batch object.
func (n *TerminalDevicePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevicePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice {
	t.Helper()
	c := &oc.CollectionTerminalDevice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevicePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice) bool) *oc.TerminalDeviceWatcher {
	t.Helper()
	w := &oc.TerminalDeviceWatcher{}
	structs := map[string]*oc.TerminalDevice{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevicePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice) bool) *oc.TerminalDeviceWatcher {
	t.Helper()
	return watch_TerminalDevicePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device to the batch object.
func (n *TerminalDevicePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_ChannelPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_ChannelPath) Get(t testing.TB) *oc.TerminalDevice_Channel {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_ChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription.
func (n *TerminalDevice_ChannelPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_ChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel)))
		return false
	})
	return c
}

func watch_TerminalDevice_ChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel) bool) *oc.TerminalDevice_ChannelWatcher {
	t.Helper()
	w := &oc.TerminalDevice_ChannelWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_ChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel) bool) *oc.TerminalDevice_ChannelWatcher {
	t.Helper()
	return watch_TerminalDevice_ChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_ChannelPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel) *oc.QualifiedTerminalDevice_Channel {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel to the batch object.
func (n *TerminalDevice_ChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_ChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_ChannelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel) bool) *oc.TerminalDevice_ChannelWatcher {
	t.Helper()
	w := &oc.TerminalDevice_ChannelWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_ChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel) bool) *oc.TerminalDevice_ChannelWatcher {
	t.Helper()
	return watch_TerminalDevice_ChannelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel to the batch object.
func (n *TerminalDevice_ChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_AdminStatePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_AdminStatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_AdminStatePath) Get(t testing.TB) oc.E_TransportTypes_AdminStateType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_AdminStatePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_AdminStateType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_AdminStatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a ONCE subscription.
func (n *TerminalDevice_Channel_AdminStatePathAny) Get(t testing.TB) []oc.E_TransportTypes_AdminStateType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_AdminStateType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_AdminStatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_AdminStateType {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_AdminStateType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_AdminStateType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_AdminStatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_AdminStateType) bool) *oc.E_TransportTypes_AdminStateTypeWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_AdminStateTypeWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_AdminStatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_AdminStateType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_AdminStatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_AdminStateType) bool) *oc.E_TransportTypes_AdminStateTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_AdminStatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_AdminStatePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_AdminStateType) *oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_AdminStateType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state to the batch object.
func (n *TerminalDevice_Channel_AdminStatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_AdminStatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_AdminStateType {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_AdminStateType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_AdminStateType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_AdminStatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_AdminStateType) bool) *oc.E_TransportTypes_AdminStateTypeWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_AdminStateTypeWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_AdminStatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_AdminStateType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_AdminStatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_AdminStateType) bool) *oc.E_TransportTypes_AdminStateTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_AdminStatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/admin-state to the batch object.
func (n *TerminalDevice_Channel_AdminStatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_AdminStatePath extracts the value of the leaf AdminState from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_AdminStateType.
func convertTerminalDevice_Channel_AdminStatePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_AdminStateType{
		Metadata: md,
	}
	val := parent.AdminState
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_AssignmentPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Assignment {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Assignment{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_AssignmentPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Assignment {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_AssignmentPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Assignment {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Assignment
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Assignment{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription.
func (n *TerminalDevice_Channel_AssignmentPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Assignment {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Assignment
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_AssignmentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Assignment {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Assignment{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Assignment) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Assignment{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Assignment)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_AssignmentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Assignment) bool) *oc.TerminalDevice_Channel_AssignmentWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_AssignmentWatcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Assignment{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Assignment)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_AssignmentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Assignment) bool) *oc.TerminalDevice_Channel_AssignmentWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_AssignmentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_AssignmentPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedTerminalDevice_Channel_Assignment {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Assignment) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment to the batch object.
func (n *TerminalDevice_Channel_AssignmentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_AssignmentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Assignment {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Assignment{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Assignment) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_AssignmentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Assignment) bool) *oc.TerminalDevice_Channel_AssignmentWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_AssignmentWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Assignment{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Assignment)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_AssignmentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Assignment) bool) *oc.TerminalDevice_Channel_AssignmentWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_AssignmentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment to the batch object.
func (n *TerminalDevice_Channel_AssignmentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_AllocationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_AllocationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_AllocationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_AllocationPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_AllocationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation to the batch object.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_AllocationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_AllocationPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_AllocationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/allocation to the batch object.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_AllocationPath extracts the value of the leaf Allocation from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Assignment_AllocationPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Allocation
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Lookup(t testing.TB) *oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Get(t testing.TB) oc.E_Assignment_AssignmentType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Assignment_AssignmentType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Get(t testing.TB) []oc.E_Assignment_AssignmentType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Assignment_AssignmentType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Assignment_AssignmentType {
	t.Helper()
	c := &oc.CollectionE_Assignment_AssignmentType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Assignment_AssignmentType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_AssignmentTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Assignment_AssignmentType) bool) *oc.E_Assignment_AssignmentTypeWatcher {
	t.Helper()
	w := &oc.E_Assignment_AssignmentTypeWatcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Assignment_AssignmentType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Assignment_AssignmentType) bool) *oc.E_Assignment_AssignmentTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_AssignmentTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Assignment_AssignmentType) *oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Assignment_AssignmentType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type to the batch object.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Assignment_AssignmentType {
	t.Helper()
	c := &oc.CollectionE_Assignment_AssignmentType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Assignment_AssignmentType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_AssignmentTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Assignment_AssignmentType) bool) *oc.E_Assignment_AssignmentTypeWatcher {
	t.Helper()
	w := &oc.E_Assignment_AssignmentTypeWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Assignment_AssignmentType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Assignment_AssignmentType) bool) *oc.E_Assignment_AssignmentTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_AssignmentTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/assignment-type to the batch object.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_AssignmentTypePath extracts the value of the leaf AssignmentType from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Assignment_AssignmentType.
func convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	qv := &oc.QualifiedE_Assignment_AssignmentType{
		Metadata: md,
	}
	val := parent.AssignmentType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_DescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description to the batch object.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_DescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_DescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_DescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/description to the batch object.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_DescriptionPath extracts the value of the leaf Description from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Assignment_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_IndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_IndexPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_IndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index to the batch object.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_IndexPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_IndexPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_IndexPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/index to the batch object.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_IndexPath extracts the value of the leaf Index from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_Assignment_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_LogicalChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_LogicalChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel to the batch object.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_LogicalChannelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_LogicalChannelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/logical-channel to the batch object.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_LogicalChannelPath extracts the value of the leaf LogicalChannel from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.LogicalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_MappingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Get(t testing.TB) oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_MappingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Get(t testing.TB) []oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_FRAME_MAPPING_PROTOCOL{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_MappingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool) *oc.E_TransportTypes_FRAME_MAPPING_PROTOCOLWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_FRAME_MAPPING_PROTOCOLWatcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_MappingPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool) *oc.E_TransportTypes_FRAME_MAPPING_PROTOCOLWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_MappingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL) *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping to the batch object.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_FRAME_MAPPING_PROTOCOL{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_MappingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool) *oc.E_TransportTypes_FRAME_MAPPING_PROTOCOLWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_FRAME_MAPPING_PROTOCOLWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_MappingPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL) bool) *oc.E_TransportTypes_FRAME_MAPPING_PROTOCOLWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_MappingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/mapping to the batch object.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_MappingPath extracts the value of the leaf Mapping from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL.
func convertTerminalDevice_Channel_Assignment_MappingPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL{
		Metadata: md,
	}
	val := parent.Mapping
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_OpticalChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_OpticalChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel to the batch object.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_OpticalChannelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_OpticalChannelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/optical-channel to the batch object.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_OpticalChannelPath extracts the value of the leaf OpticalChannel from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OpticalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_TributarySlotIndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	gs := &oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Assignment", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Await(t testing.TB, timeout time.Duration, val int32) *oc.QualifiedInt32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index to the batch object.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Assignment{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Assignment{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Assignment", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/state/tributary-slot-index to the batch object.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath extracts the value of the leaf TributarySlotIndex from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.TributarySlotIndex
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_ClientMappingModePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_ClientMappingModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_ClientMappingModePath) Get(t testing.TB) oc.E_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_ClientMappingModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a ONCE subscription.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Get(t testing.TB) []oc.E_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_CLIENT_MAPPING_MODE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_ClientMappingModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_CLIENT_MAPPING_MODE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_ClientMappingModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool) *oc.E_TransportTypes_CLIENT_MAPPING_MODEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_CLIENT_MAPPING_MODEWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_ClientMappingModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_ClientMappingModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool) *oc.E_TransportTypes_CLIENT_MAPPING_MODEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_ClientMappingModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_ClientMappingModePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_CLIENT_MAPPING_MODE) *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode to the batch object.
func (n *TerminalDevice_Channel_ClientMappingModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_CLIENT_MAPPING_MODE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_ClientMappingModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool) *oc.E_TransportTypes_CLIENT_MAPPING_MODEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_CLIENT_MAPPING_MODEWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_ClientMappingModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE) bool) *oc.E_TransportTypes_CLIENT_MAPPING_MODEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_ClientMappingModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/client-mapping-mode to the batch object.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_ClientMappingModePath extracts the value of the leaf ClientMappingMode from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE.
func convertTerminalDevice_Channel_ClientMappingModePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE{
		Metadata: md,
	}
	val := parent.ClientMappingMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a ONCE subscription.
func (n *TerminalDevice_Channel_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_DescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description to the batch object.
func (n *TerminalDevice_Channel_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_DescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_DescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_DescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/description to the batch object.
func (n *TerminalDevice_Channel_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_DescriptionPath extracts the value of the leaf Description from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_EthernetPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_EthernetPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_EthernetPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription.
func (n *TerminalDevice_Channel_EthernetPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_EthernetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Ethernet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Ethernet)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_EthernetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet) bool) *oc.TerminalDevice_Channel_EthernetWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_EthernetWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_EthernetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet) bool) *oc.TerminalDevice_Channel_EthernetWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_EthernetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_EthernetPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedTerminalDevice_Channel_Ethernet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Ethernet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet to the batch object.
func (n *TerminalDevice_Channel_EthernetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_EthernetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_EthernetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet) bool) *oc.TerminalDevice_Channel_EthernetWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_EthernetWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_EthernetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet) bool) *oc.TerminalDevice_Channel_EthernetWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_EthernetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet to the batch object.
func (n *TerminalDevice_Channel_EthernetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetAlsDelay())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_AlsDelayPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_AlsDelayPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay to the batch object.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_AlsDelayPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_AlsDelayPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/als-delay to the batch object.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_AlsDelayPath extracts the value of the leaf AlsDelay from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.AlsDelay
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Lookup(t testing.TB) *oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Ethernet_ClientAls{
		Metadata: md,
	}).SetVal(goStruct.GetClientAls())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Get(t testing.TB) oc.E_Ethernet_ClientAls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ethernet_ClientAls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Get(t testing.TB) []oc.E_Ethernet_ClientAls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ethernet_ClientAls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ethernet_ClientAls {
	t.Helper()
	c := &oc.CollectionE_Ethernet_ClientAls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ethernet_ClientAls) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_ClientAlsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ethernet_ClientAls) bool) *oc.E_Ethernet_ClientAlsWatcher {
	t.Helper()
	w := &oc.E_Ethernet_ClientAlsWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ethernet_ClientAls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ethernet_ClientAls) bool) *oc.E_Ethernet_ClientAlsWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_ClientAlsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Await(t testing.TB, timeout time.Duration, val oc.E_Ethernet_ClientAls) *oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ethernet_ClientAls) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als to the batch object.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ethernet_ClientAls {
	t.Helper()
	c := &oc.CollectionE_Ethernet_ClientAls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ethernet_ClientAls) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_ClientAlsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ethernet_ClientAls) bool) *oc.E_Ethernet_ClientAlsWatcher {
	t.Helper()
	w := &oc.E_Ethernet_ClientAlsWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ethernet_ClientAls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ethernet_ClientAls) bool) *oc.E_Ethernet_ClientAlsWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_ClientAlsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/client-als to the batch object.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_ClientAlsPath extracts the value of the leaf ClientAls from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ethernet_ClientAls.
func convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	qv := &oc.QualifiedE_Ethernet_ClientAls{
		Metadata: md,
	}
	val := parent.ClientAls
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_EsnrPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Esnr{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Esnr", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_EsnrPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_EsnrPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Esnr", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_EsnrPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet_Esnr
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_EsnrPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_Esnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Ethernet_Esnr)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_EsnrPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool) *oc.TerminalDevice_Channel_Ethernet_EsnrWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_EsnrWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Esnr", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_EsnrPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool) *oc.TerminalDevice_Channel_Ethernet_EsnrWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_EsnrPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_EsnrPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Ethernet_Esnr) *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr to the batch object.
func (n *TerminalDevice_Channel_Ethernet_EsnrPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_EsnrPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_Esnr {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_Esnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_EsnrPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool) *oc.TerminalDevice_Channel_Ethernet_EsnrWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_EsnrWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Esnr", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_EsnrPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Esnr) bool) *oc.TerminalDevice_Channel_Ethernet_EsnrWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_EsnrPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/esnr to the batch object.
func (n *TerminalDevice_Channel_Ethernet_EsnrPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
