package system

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

// Lookup fetches the value at /openconfig-system/system/mount-points/mount-point/state/size with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MountPoint_SizePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_MountPoint{}
	md, ok := oc.Lookup(t, n, "System_MountPoint", goStruct, true, false)
	if ok {
		return convertSystem_MountPoint_SizePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/mount-points/mount-point/state/size with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MountPoint_SizePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/mount-points/mount-point/state/size with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MountPoint_SizePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_MountPoint{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_MountPoint", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_MountPoint_SizePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/mount-points/mount-point/state/size with a ONCE subscription.
func (n *System_MountPoint_SizePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_SizePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_MountPoint_SizePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_MountPoint{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_MountPoint", gs, queryPath, true, false)
		return convertSystem_MountPoint_SizePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_SizePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_MountPoint_SizePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/mount-points/mount-point/state/size with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MountPoint_SizePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/mount-points/mount-point/state/size failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/size to the batch object.
func (n *System_MountPoint_SizePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_SizePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_SizePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_MountPoint_SizePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/size to the batch object.
func (n *System_MountPoint_SizePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_MountPoint_SizePath extracts the value of the leaf Size from its parent oc.System_MountPoint
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_MountPoint_SizePath(t testing.TB, md *genutil.Metadata, parent *oc.System_MountPoint) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Size
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/mount-points/mount-point/state/storage-component with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MountPoint_StorageComponentPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_MountPoint{}
	md, ok := oc.Lookup(t, n, "System_MountPoint", goStruct, true, false)
	if ok {
		return convertSystem_MountPoint_StorageComponentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/mount-points/mount-point/state/storage-component with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MountPoint_StorageComponentPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MountPoint_StorageComponentPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_MountPoint{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_MountPoint", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_MountPoint_StorageComponentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a ONCE subscription.
func (n *System_MountPoint_StorageComponentPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_StorageComponentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_MountPoint_StorageComponentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_MountPoint{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_MountPoint", gs, queryPath, true, false)
		return convertSystem_MountPoint_StorageComponentPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_StorageComponentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_MountPoint_StorageComponentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MountPoint_StorageComponentPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/mount-points/mount-point/state/storage-component failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/storage-component to the batch object.
func (n *System_MountPoint_StorageComponentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_StorageComponentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/storage-component with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_StorageComponentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_MountPoint_StorageComponentPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/storage-component to the batch object.
func (n *System_MountPoint_StorageComponentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_MountPoint_StorageComponentPath extracts the value of the leaf StorageComponent from its parent oc.System_MountPoint
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_MountPoint_StorageComponentPath(t testing.TB, md *genutil.Metadata, parent *oc.System_MountPoint) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.StorageComponent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/mount-points/mount-point/state/utilized with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MountPoint_UtilizedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_MountPoint{}
	md, ok := oc.Lookup(t, n, "System_MountPoint", goStruct, true, false)
	if ok {
		return convertSystem_MountPoint_UtilizedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/mount-points/mount-point/state/utilized with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MountPoint_UtilizedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/mount-points/mount-point/state/utilized with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MountPoint_UtilizedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_MountPoint{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_MountPoint", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_MountPoint_UtilizedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/mount-points/mount-point/state/utilized with a ONCE subscription.
func (n *System_MountPoint_UtilizedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/utilized with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_UtilizedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_MountPoint_UtilizedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_MountPoint{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_MountPoint", gs, queryPath, true, false)
		return convertSystem_MountPoint_UtilizedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/utilized with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_UtilizedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_MountPoint_UtilizedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/mount-points/mount-point/state/utilized with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MountPoint_UtilizedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/mount-points/mount-point/state/utilized failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/utilized to the batch object.
func (n *System_MountPoint_UtilizedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/utilized with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_UtilizedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/utilized with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_UtilizedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_MountPoint_UtilizedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/utilized to the batch object.
func (n *System_MountPoint_UtilizedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_MountPoint_UtilizedPath extracts the value of the leaf Utilized from its parent oc.System_MountPoint
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_MountPoint_UtilizedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_MountPoint) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Utilized
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_NtpPath) Lookup(t testing.TB) *oc.QualifiedSystem_Ntp {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Ntp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_NtpPath) Get(t testing.TB) *oc.System_Ntp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_NtpPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Ntp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Ntp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Ntp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp with a ONCE subscription.
func (n *System_NtpPathAny) Get(t testing.TB) []*oc.System_Ntp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Ntp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_NtpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Ntp {
	t.Helper()
	c := &oc.CollectionSystem_Ntp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Ntp) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Ntp{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Ntp)))
		return false
	})
	return c
}

func watch_System_NtpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Ntp) bool) *oc.System_NtpWatcher {
	t.Helper()
	w := &oc.System_NtpWatcher{}
	gs := &oc.System_Ntp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Ntp{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Ntp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_NtpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Ntp) bool) *oc.System_NtpWatcher {
	t.Helper()
	return watch_System_NtpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_NtpPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Ntp) *oc.QualifiedSystem_Ntp {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Ntp) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp to the batch object.
func (n *System_NtpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_NtpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Ntp {
	t.Helper()
	c := &oc.CollectionSystem_Ntp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Ntp) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_NtpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Ntp) bool) *oc.System_NtpWatcher {
	t.Helper()
	return watch_System_NtpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp to the batch object.
func (n *System_NtpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/ntp/state/auth-mismatch with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_AuthMismatchPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_AuthMismatchPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/state/auth-mismatch with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_AuthMismatchPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/state/auth-mismatch with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_AuthMismatchPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_AuthMismatchPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/state/auth-mismatch with a ONCE subscription.
func (n *System_Ntp_AuthMismatchPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/auth-mismatch with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_AuthMismatchPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_AuthMismatchPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Ntp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp", gs, queryPath, true, false)
		return convertSystem_Ntp_AuthMismatchPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/auth-mismatch with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_AuthMismatchPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Ntp_AuthMismatchPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/state/auth-mismatch with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_AuthMismatchPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/state/auth-mismatch failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/state/auth-mismatch to the batch object.
func (n *System_Ntp_AuthMismatchPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/auth-mismatch with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_AuthMismatchPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/auth-mismatch with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_AuthMismatchPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Ntp_AuthMismatchPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/state/auth-mismatch to the batch object.
func (n *System_Ntp_AuthMismatchPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_AuthMismatchPath extracts the value of the leaf AuthMismatch from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Ntp_AuthMismatchPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AuthMismatch
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/state/enable-ntp-auth with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_EnableNtpAuthPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_EnableNtpAuthPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableNtpAuth())
}

// Get fetches the value at /openconfig-system/system/ntp/state/enable-ntp-auth with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_EnableNtpAuthPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/state/enable-ntp-auth with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_EnableNtpAuthPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_EnableNtpAuthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/state/enable-ntp-auth with a ONCE subscription.
func (n *System_Ntp_EnableNtpAuthPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/enable-ntp-auth with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_EnableNtpAuthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_EnableNtpAuthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_Ntp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp", gs, queryPath, true, false)
		return convertSystem_Ntp_EnableNtpAuthPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/enable-ntp-auth with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_EnableNtpAuthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_Ntp_EnableNtpAuthPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/state/enable-ntp-auth with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_EnableNtpAuthPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/state/enable-ntp-auth failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/state/enable-ntp-auth to the batch object.
func (n *System_Ntp_EnableNtpAuthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/enable-ntp-auth with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_EnableNtpAuthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/enable-ntp-auth with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_EnableNtpAuthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_Ntp_EnableNtpAuthPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/state/enable-ntp-auth to the batch object.
func (n *System_Ntp_EnableNtpAuthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_EnableNtpAuthPath extracts the value of the leaf EnableNtpAuth from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Ntp_EnableNtpAuthPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableNtpAuth
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-system/system/ntp/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/state/enabled with a ONCE subscription.
func (n *System_Ntp_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_Ntp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp", gs, queryPath, true, false)
		return convertSystem_Ntp_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_Ntp_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/state/enabled to the batch object.
func (n *System_Ntp_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_Ntp_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/state/enabled to the batch object.
func (n *System_Ntp_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_EnabledPath extracts the value of the leaf Enabled from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Ntp_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKeyPath) Lookup(t testing.TB) *oc.QualifiedSystem_Ntp_NtpKey {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Ntp_NtpKey{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKeyPath) Get(t testing.TB) *oc.System_Ntp_NtpKey {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Ntp_NtpKey {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Ntp_NtpKey
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Ntp_NtpKey{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription.
func (n *System_Ntp_NtpKeyPathAny) Get(t testing.TB) []*oc.System_Ntp_NtpKey {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Ntp_NtpKey
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKeyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Ntp_NtpKey {
	t.Helper()
	c := &oc.CollectionSystem_Ntp_NtpKey{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Ntp_NtpKey) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Ntp_NtpKey{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Ntp_NtpKey)))
		return false
	})
	return c
}

func watch_System_Ntp_NtpKeyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Ntp_NtpKey) bool) *oc.System_Ntp_NtpKeyWatcher {
	t.Helper()
	w := &oc.System_Ntp_NtpKeyWatcher{}
	gs := &oc.System_Ntp_NtpKey{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp_NtpKey", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Ntp_NtpKey{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Ntp_NtpKey)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKeyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Ntp_NtpKey) bool) *oc.System_Ntp_NtpKeyWatcher {
	t.Helper()
	return watch_System_Ntp_NtpKeyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_NtpKeyPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Ntp_NtpKey) *oc.QualifiedSystem_Ntp_NtpKey {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Ntp_NtpKey) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/ntp-keys/ntp-key failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key to the batch object.
func (n *System_Ntp_NtpKeyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKeyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Ntp_NtpKey {
	t.Helper()
	c := &oc.CollectionSystem_Ntp_NtpKey{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Ntp_NtpKey) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKeyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Ntp_NtpKey) bool) *oc.System_Ntp_NtpKeyWatcher {
	t.Helper()
	return watch_System_Ntp_NtpKeyPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key to the batch object.
func (n *System_Ntp_NtpKeyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKey_KeyIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_NtpKey_KeyIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Ntp_NtpKey{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp_NtpKey", gs, queryPath, true, false)
		return convertSystem_Ntp_NtpKey_KeyIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKey_KeyIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Ntp_NtpKey_KeyIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_NtpKey_KeyIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id to the batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Ntp_NtpKey_KeyIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-id to the batch object.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_NtpKey_KeyIdPath extracts the value of the leaf KeyId from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Ntp_NtpKey_KeyIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.KeyId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyTypePath) Lookup(t testing.TB) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyTypePath) Get(t testing.TB) oc.E_System_NTP_AUTH_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_System_NTP_AUTH_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Get(t testing.TB) []oc.E_System_NTP_AUTH_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_System_NTP_AUTH_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKey_KeyTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_System_NTP_AUTH_TYPE {
	t.Helper()
	c := &oc.CollectionE_System_NTP_AUTH_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_System_NTP_AUTH_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_NtpKey_KeyTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_System_NTP_AUTH_TYPE) bool) *oc.E_System_NTP_AUTH_TYPEWatcher {
	t.Helper()
	w := &oc.E_System_NTP_AUTH_TYPEWatcher{}
	gs := &oc.System_Ntp_NtpKey{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp_NtpKey", gs, queryPath, true, false)
		return convertSystem_Ntp_NtpKey_KeyTypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_System_NTP_AUTH_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKey_KeyTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_System_NTP_AUTH_TYPE) bool) *oc.E_System_NTP_AUTH_TYPEWatcher {
	t.Helper()
	return watch_System_Ntp_NtpKey_KeyTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_NtpKey_KeyTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_System_NTP_AUTH_TYPE) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_System_NTP_AUTH_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type to the batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_System_NTP_AUTH_TYPE {
	t.Helper()
	c := &oc.CollectionE_System_NTP_AUTH_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_System_NTP_AUTH_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_System_NTP_AUTH_TYPE) bool) *oc.E_System_NTP_AUTH_TYPEWatcher {
	t.Helper()
	return watch_System_Ntp_NtpKey_KeyTypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-type to the batch object.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_NtpKey_KeyTypePath extracts the value of the leaf KeyType from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedE_System_NTP_AUTH_TYPE.
func convertSystem_Ntp_NtpKey_KeyTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_System_NTP_AUTH_TYPE{
		Metadata: md,
	}
	val := parent.KeyType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyValuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyValuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKey_KeyValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_NtpKey_KeyValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Ntp_NtpKey{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp_NtpKey", gs, queryPath, true, false)
		return convertSystem_Ntp_NtpKey_KeyValuePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKey_KeyValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Ntp_NtpKey_KeyValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_NtpKey_KeyValuePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value to the batch object.
func (n *System_Ntp_NtpKey_KeyValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Ntp_NtpKey_KeyValuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/ntp-keys/ntp-key/state/key-value to the batch object.
func (n *System_Ntp_NtpKey_KeyValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_NtpKey_KeyValuePath extracts the value of the leaf KeyValue from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Ntp_NtpKey_KeyValuePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.KeyValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/state/ntp-source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpSourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, false)
	if ok {
		return convertSystem_Ntp_NtpSourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/state/ntp-source-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpSourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/state/ntp-source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpSourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpSourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/state/ntp-source-address with a ONCE subscription.
func (n *System_Ntp_NtpSourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/ntp-source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpSourceAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Ntp_NtpSourceAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Ntp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Ntp", gs, queryPath, true, false)
		return convertSystem_Ntp_NtpSourceAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/ntp-source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpSourceAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Ntp_NtpSourceAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ntp/state/ntp-source-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Ntp_NtpSourceAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ntp/state/ntp-source-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ntp/state/ntp-source-address to the batch object.
func (n *System_Ntp_NtpSourceAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ntp/state/ntp-source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Ntp_NtpSourceAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ntp/state/ntp-source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Ntp_NtpSourceAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Ntp_NtpSourceAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ntp/state/ntp-source-address to the batch object.
func (n *System_Ntp_NtpSourceAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Ntp_NtpSourceAddressPath extracts the value of the leaf NtpSourceAddress from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Ntp_NtpSourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NtpSourceAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
