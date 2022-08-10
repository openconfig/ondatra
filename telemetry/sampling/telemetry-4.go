package sampling

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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/state/enabled with a ONCE subscription.
func (n *Sampling_Sflow_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Sampling_Sflow_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/state/enabled to the batch object.
func (n *Sampling_Sflow_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Sampling_Sflow{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling_Sflow", structs[pre], queryPath, true, false)
			qv := convertSampling_Sflow_EnabledPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Sampling_Sflow_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/state/enabled to the batch object.
func (n *Sampling_Sflow_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_EnabledPath extracts the value of the leaf Enabled from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSampling_Sflow_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_IngressSamplingRatePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_IngressSamplingRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_IngressSamplingRatePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_IngressSamplingRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a ONCE subscription.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_IngressSamplingRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_IngressSamplingRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_IngressSamplingRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_IngressSamplingRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Sampling_Sflow_IngressSamplingRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_IngressSamplingRatePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate to the batch object.
func (n *Sampling_Sflow_IngressSamplingRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_IngressSamplingRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Sampling_Sflow{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling_Sflow", structs[pre], queryPath, true, false)
			qv := convertSampling_Sflow_IngressSamplingRatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Sampling_Sflow_IngressSamplingRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/state/ingress-sampling-rate to the batch object.
func (n *Sampling_Sflow_IngressSamplingRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_IngressSamplingRatePath extracts the value of the leaf IngressSamplingRate from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertSampling_Sflow_IngressSamplingRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.IngressSamplingRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_InterfacePath) Lookup(t testing.TB) *oc.QualifiedSampling_Sflow_Interface {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Interface{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSampling_Sflow_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_InterfacePath) Get(t testing.TB) *oc.Sampling_Sflow_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedSampling_Sflow_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling_Sflow_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling_Sflow_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a ONCE subscription.
func (n *Sampling_Sflow_InterfacePathAny) Get(t testing.TB) []*oc.Sampling_Sflow_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling_Sflow_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling_Sflow_Interface {
	t.Helper()
	c := &oc.CollectionSampling_Sflow_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling_Sflow_Interface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSampling_Sflow_Interface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Sampling_Sflow_Interface)))
		return false
	})
	return c
}

func watch_Sampling_Sflow_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Interface) bool) *oc.Sampling_Sflow_InterfaceWatcher {
	t.Helper()
	w := &oc.Sampling_Sflow_InterfaceWatcher{}
	gs := &oc.Sampling_Sflow_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow_Interface", gs, queryPath, false, false)
		qv := (&oc.QualifiedSampling_Sflow_Interface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling_Sflow_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Interface) bool) *oc.Sampling_Sflow_InterfaceWatcher {
	t.Helper()
	return watch_Sampling_Sflow_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Sampling_Sflow_Interface) *oc.QualifiedSampling_Sflow_Interface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSampling_Sflow_Interface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/interfaces/interface to the batch object.
func (n *Sampling_Sflow_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling_Sflow_Interface {
	t.Helper()
	c := &oc.CollectionSampling_Sflow_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling_Sflow_Interface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Interface) bool) *oc.Sampling_Sflow_InterfaceWatcher {
	t.Helper()
	w := &oc.Sampling_Sflow_InterfaceWatcher{}
	structs := map[string]*oc.Sampling_Sflow_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Sampling_Sflow_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling_Sflow_Interface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSampling_Sflow_Interface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling_Sflow_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Interface) bool) *oc.Sampling_Sflow_InterfaceWatcher {
	t.Helper()
	return watch_Sampling_Sflow_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/interfaces/interface to the batch object.
func (n *Sampling_Sflow_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
