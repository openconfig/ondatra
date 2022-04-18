package device

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

// Lookup fetches the value at / with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *DevicePath) Lookup(t testing.TB) *oc.QualifiedDevice {
	t.Helper()
	goStruct := &oc.Device{}
	md, ok := oc.Lookup(t, n, "Device", goStruct, false, false)
	if ok {
		return (&oc.QualifiedDevice{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at / with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *DevicePath) Get(t testing.TB) *oc.Device {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Collect starts an asynchronous collection of the values at / with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *DevicePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionDevice {
	t.Helper()
	c := &oc.CollectionDevice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedDevice) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedDevice{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Device)))
		return false
	})
	return c
}

func watch_DevicePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedDevice) bool) *oc.DeviceWatcher {
	t.Helper()
	w := &oc.DeviceWatcher{}
	gs := &oc.Device{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Device", gs, queryPath, false, false)
		qv := (&oc.QualifiedDevice{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedDevice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at / with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *DevicePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedDevice) bool) *oc.DeviceWatcher {
	t.Helper()
	return watch_DevicePath(t, n, timeout, predicate)
}

// Await observes values at / with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *DevicePath) Await(t testing.TB, timeout time.Duration, val *oc.Device) *oc.QualifiedDevice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedDevice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at / failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds / to the batch object.
func (n *DevicePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
