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

// Lookup fetches the value at /openconfig-system/system/license with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LicensePath) Lookup(t testing.TB) *oc.QualifiedSystem_License {
	t.Helper()
	goStruct := &oc.System_License{}
	md, ok := oc.Lookup(t, n, "System_License", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LicensePath) Get(t testing.TB) *oc.System_License {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LicensePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license with a ONCE subscription.
func (n *System_LicensePathAny) Get(t testing.TB) []*oc.System_License {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_License
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LicensePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License {
	t.Helper()
	c := &oc.CollectionSystem_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_License{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_License)))
		return false
	})
	return c
}

func watch_System_LicensePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	w := &oc.System_LicenseWatcher{}
	gs := &oc.System_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LicensePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	return watch_System_LicensePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_LicensePath) Await(t testing.TB, timeout time.Duration, val *oc.System_License) *oc.QualifiedSystem_License {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_License) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license to the batch object.
func (n *System_LicensePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LicensePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License {
	t.Helper()
	c := &oc.CollectionSystem_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LicensePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	return watch_System_LicensePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license to the batch object.
func (n *System_LicensePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_LicensePath) Lookup(t testing.TB) *oc.QualifiedSystem_License_License {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_LicensePath) Get(t testing.TB) *oc.System_License_License {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_LicensePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License_License {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License_License
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license with a ONCE subscription.
func (n *System_License_LicensePathAny) Get(t testing.TB) []*oc.System_License_License {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_License_License
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_LicensePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License {
	t.Helper()
	c := &oc.CollectionSystem_License_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_License_License{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_License_License)))
		return false
	})
	return c
}

func watch_System_License_LicensePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	w := &oc.System_License_LicenseWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License_License)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_LicensePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	return watch_System_License_LicensePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_LicensePath) Await(t testing.TB, timeout time.Duration, val *oc.System_License_License) *oc.QualifiedSystem_License_License {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_License_License) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license to the batch object.
func (n *System_License_LicensePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_LicensePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License {
	t.Helper()
	c := &oc.CollectionSystem_License_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_LicensePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	return watch_System_License_LicensePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license to the batch object.
func (n *System_License_LicensePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ActivePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ActivePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetActive())
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ActivePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ActivePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ActivePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription.
func (n *System_License_License_ActivePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ActivePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ActivePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_ActivePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ActivePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ActivePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ActivePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/active failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/active to the batch object.
func (n *System_License_License_ActivePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ActivePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ActivePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ActivePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/active to the batch object.
func (n *System_License_License_ActivePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ActivePath extracts the value of the leaf Active from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ActivePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Active
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription.
func (n *System_License_License_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_DescriptionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/description to the batch object.
func (n *System_License_License_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_DescriptionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/description to the batch object.
func (n *System_License_License_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_DescriptionPath extracts the value of the leaf Description from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_License_License_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ExpirationDatePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ExpirationDatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ExpirationDatePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ExpirationDatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ExpirationDatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription.
func (n *System_License_License_ExpirationDatePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpirationDatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ExpirationDatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_ExpirationDatePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpirationDatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_ExpirationDatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ExpirationDatePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/expiration-date failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expiration-date to the batch object.
func (n *System_License_License_ExpirationDatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpirationDatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpirationDatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_ExpirationDatePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expiration-date to the batch object.
func (n *System_License_License_ExpirationDatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ExpirationDatePath extracts the value of the leaf ExpirationDate from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_License_License_ExpirationDatePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ExpirationDate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ExpiredPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ExpiredPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ExpiredPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ExpiredPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ExpiredPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription.
func (n *System_License_License_ExpiredPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpiredPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ExpiredPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_ExpiredPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpiredPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ExpiredPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ExpiredPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/expired failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expired to the batch object.
func (n *System_License_License_ExpiredPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpiredPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpiredPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ExpiredPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expired to the batch object.
func (n *System_License_License_ExpiredPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ExpiredPath extracts the value of the leaf Expired from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ExpiredPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Expired
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_InUsePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_InUsePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_InUsePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_InUsePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_InUsePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription.
func (n *System_License_License_InUsePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_InUsePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_InUsePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_InUsePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_InUsePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_InUsePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_InUsePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/in-use failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/in-use to the batch object.
func (n *System_License_License_InUsePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_InUsePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_InUsePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_InUsePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/in-use to the batch object.
func (n *System_License_License_InUsePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_InUsePath extracts the value of the leaf InUse from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_InUsePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.InUse
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_IssueDatePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_IssueDatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_IssueDatePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_IssueDatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_IssueDatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription.
func (n *System_License_License_IssueDatePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_IssueDatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_IssueDatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_IssueDatePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_IssueDatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_IssueDatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_IssueDatePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/issue-date failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/issue-date to the batch object.
func (n *System_License_License_IssueDatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_IssueDatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_IssueDatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_IssueDatePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/issue-date to the batch object.
func (n *System_License_License_IssueDatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_IssueDatePath extracts the value of the leaf IssueDate from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_License_License_IssueDatePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.IssueDate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_LicenseDataPath) Lookup(t testing.TB) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_LicenseDataPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_LicenseDataPath) Get(t testing.TB) oc.System_License_License_LicenseData_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_LicenseDataPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License_License_LicenseData_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_LicenseDataPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription.
func (n *System_License_License_LicenseDataPathAny) Get(t testing.TB) []oc.System_License_License_LicenseData_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_License_License_LicenseData_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseDataPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License_LicenseData_Union {
	t.Helper()
	c := &oc.CollectionSystem_License_License_LicenseData_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License_LicenseData_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_LicenseDataPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	w := &oc.System_License_License_LicenseData_UnionWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_LicenseDataPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License_License_LicenseData_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseDataPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	return watch_System_License_License_LicenseDataPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_LicenseDataPath) Await(t testing.TB, timeout time.Duration, val oc.System_License_License_LicenseData_Union) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_License_License_LicenseData_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/license-data failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-data to the batch object.
func (n *System_License_License_LicenseDataPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseDataPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License_LicenseData_Union {
	t.Helper()
	c := &oc.CollectionSystem_License_License_LicenseData_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License_LicenseData_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseDataPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	return watch_System_License_License_LicenseDataPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-data to the batch object.
func (n *System_License_License_LicenseDataPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_LicenseDataPath extracts the value of the leaf LicenseData from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_License_License_LicenseData_Union.
func convertSystem_License_License_LicenseDataPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_License_License_LicenseData_Union{
		Metadata: md,
	}
	val := parent.LicenseData
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_LicenseIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_LicenseIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_LicenseIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_LicenseIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_LicenseIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription.
func (n *System_License_License_LicenseIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_LicenseIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_LicenseIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_LicenseIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_LicenseIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/license-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-id to the batch object.
func (n *System_License_License_LicenseIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_LicenseIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-id to the batch object.
func (n *System_License_License_LicenseIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_LicenseIdPath extracts the value of the leaf LicenseId from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_License_License_LicenseIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LicenseId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ValidPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ValidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ValidPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ValidPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ValidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription.
func (n *System_License_License_ValidPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ValidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ValidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return convertSystem_License_License_ValidPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ValidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ValidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ValidPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/valid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/valid to the batch object.
func (n *System_License_License_ValidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ValidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ValidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ValidPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/valid to the batch object.
func (n *System_License_License_ValidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ValidPath extracts the value of the leaf Valid from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ValidPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Valid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LoggingPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging {
	t.Helper()
	goStruct := &oc.System_Logging{}
	md, ok := oc.Lookup(t, n, "System_Logging", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LoggingPath) Get(t testing.TB) *oc.System_Logging {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LoggingPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging with a ONCE subscription.
func (n *System_LoggingPathAny) Get(t testing.TB) []*oc.System_Logging {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LoggingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging {
	t.Helper()
	c := &oc.CollectionSystem_Logging{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Logging{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Logging)))
		return false
	})
	return c
}

func watch_System_LoggingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	w := &oc.System_LoggingWatcher{}
	gs := &oc.System_Logging{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LoggingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	return watch_System_LoggingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_LoggingPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Logging) *oc.QualifiedSystem_Logging {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Logging) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging to the batch object.
func (n *System_LoggingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LoggingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging {
	t.Helper()
	c := &oc.CollectionSystem_Logging{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LoggingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	return watch_System_LoggingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging to the batch object.
func (n *System_LoggingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
