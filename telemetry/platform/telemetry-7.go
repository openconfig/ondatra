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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_NonFatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_NonFatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/tlp-prefix-blocked-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath extracts the value of the leaf TlpPrefixBlockedErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_TlpPrefixBlockedErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TlpPrefixBlockedErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_TotalErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_NonFatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_NonFatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_NonFatalErrors_TotalErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_TotalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_TotalErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_TotalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_TotalErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/total-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_TotalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_TotalErrorsPath extracts the value of the leaf TotalErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_TotalErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TotalErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_UndefinedErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_NonFatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_NonFatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_NonFatalErrors_UndefinedErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_UndefinedErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_UndefinedErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_UndefinedErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_UndefinedErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/undefined-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_UndefinedErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_UndefinedErrorsPath extracts the value of the leaf UndefinedErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_UndefinedErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UndefinedErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_NonFatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_NonFatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/unexpected-completion-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath extracts the value of the leaf UnexpectedCompletionErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_UnexpectedCompletionErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UnexpectedCompletionErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_NonFatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_NonFatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/unsupported-request-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_UnsupportedRequestErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath extracts the value of the leaf UnsupportedRequestErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_UnsupportedRequestErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UnsupportedRequestErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PortPath) Lookup(t testing.TB) *oc.QualifiedComponent_Port {
	t.Helper()
	goStruct := &oc.Component_Port{}
	md, ok := oc.Lookup(t, n, "Component_Port", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Port{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PortPath) Get(t testing.TB) *oc.Component_Port {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port with a ONCE subscription.
func (n *Component_PortPathAny) Get(t testing.TB) []*oc.Component_Port {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Port {
	t.Helper()
	c := &oc.CollectionComponent_Port{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Port) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Port{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Port)))
		return false
	})
	return c
}

func watch_Component_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Port) bool) *oc.Component_PortWatcher {
	t.Helper()
	w := &oc.Component_PortWatcher{}
	gs := &oc.Component_Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Port{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Port)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Port) bool) *oc.Component_PortWatcher {
	t.Helper()
	return watch_Component_PortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_PortPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Port) *oc.QualifiedComponent_Port {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Port) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port to the batch object.
func (n *Component_PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Port {
	t.Helper()
	c := &oc.CollectionComponent_Port{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Port) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Port) bool) *oc.Component_PortWatcher {
	t.Helper()
	return watch_Component_PortPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port to the batch object.
func (n *Component_PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutModePath) Lookup(t testing.TB) *oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutModePath) Get(t testing.TB) *oc.Component_Port_BreakoutMode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutModePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port_BreakoutMode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
func (n *Component_Port_BreakoutModePathAny) Get(t testing.TB) []*oc.Component_Port_BreakoutMode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port_BreakoutMode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Port_BreakoutMode {
	t.Helper()
	c := &oc.CollectionComponent_Port_BreakoutMode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Port_BreakoutMode) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Port_BreakoutMode)))
		return false
	})
	return c
}

func watch_Component_Port_BreakoutModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Port_BreakoutMode) bool) *oc.Component_Port_BreakoutModeWatcher {
	t.Helper()
	w := &oc.Component_Port_BreakoutModeWatcher{}
	gs := &oc.Component_Port_BreakoutMode{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port_BreakoutMode", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Port_BreakoutMode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Port_BreakoutMode) bool) *oc.Component_Port_BreakoutModeWatcher {
	t.Helper()
	return watch_Component_Port_BreakoutModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port/breakout-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Port_BreakoutModePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Port_BreakoutMode) *oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Port_BreakoutMode) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port/breakout-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode to the batch object.
func (n *Component_Port_BreakoutModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Port_BreakoutMode {
	t.Helper()
	c := &oc.CollectionComponent_Port_BreakoutMode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Port_BreakoutMode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Port_BreakoutMode) bool) *oc.Component_Port_BreakoutModeWatcher {
	t.Helper()
	return watch_Component_Port_BreakoutModePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode to the batch object.
func (n *Component_Port_BreakoutModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_GroupPath) Lookup(t testing.TB) *oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_GroupPath) Get(t testing.TB) *oc.Component_Port_BreakoutMode_Group {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_GroupPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port_BreakoutMode_Group
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
func (n *Component_Port_BreakoutMode_GroupPathAny) Get(t testing.TB) []*oc.Component_Port_BreakoutMode_Group {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port_BreakoutMode_Group
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_GroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Port_BreakoutMode_Group {
	t.Helper()
	c := &oc.CollectionComponent_Port_BreakoutMode_Group{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Port_BreakoutMode_Group) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Port_BreakoutMode_Group)))
		return false
	})
	return c
}

func watch_Component_Port_BreakoutMode_GroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Port_BreakoutMode_Group) bool) *oc.Component_Port_BreakoutMode_GroupWatcher {
	t.Helper()
	w := &oc.Component_Port_BreakoutMode_GroupWatcher{}
	gs := &oc.Component_Port_BreakoutMode_Group{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port_BreakoutMode_Group", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Port_BreakoutMode_Group)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_GroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Port_BreakoutMode_Group) bool) *oc.Component_Port_BreakoutMode_GroupWatcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_GroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Port_BreakoutMode_GroupPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Port_BreakoutMode_Group) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port/breakout-mode/groups/group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group to the batch object.
func (n *Component_Port_BreakoutMode_GroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_GroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Port_BreakoutMode_Group {
	t.Helper()
	c := &oc.CollectionComponent_Port_BreakoutMode_Group{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Port_BreakoutMode_Group) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_GroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Port_BreakoutMode_Group) bool) *oc.Component_Port_BreakoutMode_GroupWatcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_GroupPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group to the batch object.
func (n *Component_Port_BreakoutMode_GroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, false)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Get(t testing.TB) oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Get(t testing.TB) []oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_ETHERNET_SPEED
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_ETHERNET_SPEED{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Port_BreakoutMode_Group_BreakoutSpeedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	w := &oc.E_IfEthernet_ETHERNET_SPEEDWatcher{}
	gs := &oc.Component_Port_BreakoutMode_Group{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port_BreakoutMode_Group", gs, queryPath, true, false)
		return convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfEthernet_ETHERNET_SPEED)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_BreakoutSpeedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfEthernet_ETHERNET_SPEED) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed to the batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c := &oc.CollectionE_IfEthernet_ETHERNET_SPEED{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfEthernet_ETHERNET_SPEED) bool) *oc.E_IfEthernet_ETHERNET_SPEEDWatcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_BreakoutSpeedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/breakout-speed to the batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath extracts the value of the leaf BreakoutSpeed from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_ETHERNET_SPEED.
func convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_ETHERNET_SPEED{
		Metadata: md,
	}
	val := parent.BreakoutSpeed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, false)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Port_BreakoutMode_Group_IndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Component_Port_BreakoutMode_Group{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port_BreakoutMode_Group", gs, queryPath, true, false)
		return convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_IndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index to the batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_IndexPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/index to the batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Port_BreakoutMode_Group_IndexPath extracts the value of the leaf Index from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, false)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Port_BreakoutMode_Group_NumBreakoutsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Component_Port_BreakoutMode_Group{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port_BreakoutMode_Group", gs, queryPath, true, false)
		return convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_NumBreakoutsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts to the batch object.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_NumBreakoutsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-breakouts to the batch object.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath extracts the value of the leaf NumBreakouts from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.NumBreakouts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
