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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrorsPath) Lookup(t testing.TB) *oc.QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrorsPath) Get(t testing.TB) *oc.Component_Pcie_FatalErrors {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Pcie_FatalErrors
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrorsPathAny) Get(t testing.TB) []*oc.Component_Pcie_FatalErrors {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Pcie_FatalErrors
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie_FatalErrors {
	t.Helper()
	c := &oc.CollectionComponent_Pcie_FatalErrors{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie_FatalErrors) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Pcie_FatalErrors)))
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	w := &oc.Component_Pcie_FatalErrorsWatcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie_FatalErrors)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrorsPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Pcie_FatalErrors) *oc.QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Pcie_FatalErrors) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors to the batch object.
func (n *Component_Pcie_FatalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie_FatalErrors {
	t.Helper()
	c := &oc.CollectionComponent_Pcie_FatalErrors{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie_FatalErrors) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors to the batch object.
func (n *Component_Pcie_FatalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_AcsViolationErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AcsViolationErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AcsViolationErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath extracts the value of the leaf AcsViolationErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AcsViolationErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath extracts the value of the leaf AtomicOpBlockedErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AtomicOpBlockedErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors to the batch object.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors to the batch object.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath extracts the value of the leaf BlockedTlpErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.BlockedTlpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors to the batch object.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors to the batch object.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath extracts the value of the leaf CompletionAbortErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CompletionAbortErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors to the batch object.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/completion-timeout-errors to the batch object.
func (n *Component_Pcie_FatalErrors_CompletionTimeoutErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_CompletionTimeoutErrorsPath extracts the value of the leaf CompletionTimeoutErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_CompletionTimeoutErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CompletionTimeoutErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_DataLinkErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_DataLinkErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_DataLinkErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_DataLinkErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_DataLinkErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors to the batch object.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_DataLinkErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/data-link-errors to the batch object.
func (n *Component_Pcie_FatalErrors_DataLinkErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_DataLinkErrorsPath extracts the value of the leaf DataLinkErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_DataLinkErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DataLinkErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_EcrcErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_EcrcErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_EcrcErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_EcrcErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_EcrcErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors to the batch object.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_EcrcErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/ecrc-errors to the batch object.
func (n *Component_Pcie_FatalErrors_EcrcErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_EcrcErrorsPath extracts the value of the leaf EcrcErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_EcrcErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.EcrcErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors to the batch object.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/flow-control-protocol-errors to the batch object.
func (n *Component_Pcie_FatalErrors_FlowControlProtocolErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_FlowControlProtocolErrorsPath extracts the value of the leaf FlowControlProtocolErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_FlowControlProtocolErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FlowControlProtocolErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_InternalErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_InternalErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_InternalErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_InternalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_InternalErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_InternalErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_InternalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_InternalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_InternalErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_InternalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_InternalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_InternalErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors to the batch object.
func (n *Component_Pcie_FatalErrors_InternalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_InternalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_InternalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_InternalErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/internal-errors to the batch object.
func (n *Component_Pcie_FatalErrors_InternalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_InternalErrorsPath extracts the value of the leaf InternalErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_InternalErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InternalErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_MalformedTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_MalformedTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_MalformedTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_FatalErrors_MalformedTlpErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_MalformedTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors to the batch object.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_MalformedTlpErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/malformed-tlp-errors to the batch object.
func (n *Component_Pcie_FatalErrors_MalformedTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_MalformedTlpErrorsPath extracts the value of the leaf MalformedTlpErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_MalformedTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MalformedTlpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
