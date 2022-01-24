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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/blocked-tlp-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_BlockedTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_BlockedTlpErrorsPath extracts the value of the leaf BlockedTlpErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_BlockedTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-abort-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_CompletionAbortErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_CompletionAbortErrorsPath extracts the value of the leaf CompletionAbortErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_CompletionAbortErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/completion-timeout-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_CompletionTimeoutErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath extracts the value of the leaf CompletionTimeoutErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_CompletionTimeoutErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_DataLinkErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_DataLinkErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_DataLinkErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_DataLinkErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_DataLinkErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_DataLinkErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/data-link-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_DataLinkErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_DataLinkErrorsPath extracts the value of the leaf DataLinkErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_DataLinkErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_EcrcErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_EcrcErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_EcrcErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_EcrcErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_EcrcErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_EcrcErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/ecrc-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_EcrcErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_EcrcErrorsPath extracts the value of the leaf EcrcErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_EcrcErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/flow-control-protocol-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_FlowControlProtocolErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath extracts the value of the leaf FlowControlProtocolErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_FlowControlProtocolErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_InternalErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_InternalErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_InternalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_InternalErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_InternalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_InternalErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/internal-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_InternalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_InternalErrorsPath extracts the value of the leaf InternalErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_InternalErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/malformed-tlp-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_MalformedTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_MalformedTlpErrorsPath extracts the value of the leaf MalformedTlpErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_MalformedTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/poisoned-tlp-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_PoisonedTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_PoisonedTlpErrorsPath extracts the value of the leaf PoisonedTlpErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_PoisonedTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.PoisonedTlpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/receiver-overflow-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_ReceiverOverflowErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath extracts the value of the leaf ReceiverOverflowErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_ReceiverOverflowErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ReceiverOverflowErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_NonFatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_NonFatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a ONCE subscription.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_NonFatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_NonFatalErrors", gs, queryPath, true, false)
		return convertComponent_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/non-fatal-errors/surprise-down-errors to the batch object.
func (n *Component_Pcie_NonFatalErrors_SurpriseDownErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_NonFatalErrors_SurpriseDownErrorsPath extracts the value of the leaf SurpriseDownErrors from its parent oc.Component_Pcie_NonFatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_NonFatalErrors_SurpriseDownErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_NonFatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.SurpriseDownErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
