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

// Lookup fetches the value at /openconfig-sampling/sampling with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SamplingPath) Lookup(t testing.TB) *oc.QualifiedSampling {
	t.Helper()
	goStruct := &oc.Sampling{}
	md, ok := oc.Lookup(t, n, "Sampling", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSampling{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SamplingPath) Get(t testing.TB) *oc.Sampling {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SamplingPathAny) Lookup(t testing.TB) []*oc.QualifiedSampling {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling with a ONCE subscription.
func (n *SamplingPathAny) Get(t testing.TB) []*oc.Sampling {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *SamplingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling {
	t.Helper()
	c := &oc.CollectionSampling{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSampling{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Sampling)))
		return false
	})
	return c
}

func watch_SamplingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling) bool) *oc.SamplingWatcher {
	t.Helper()
	w := &oc.SamplingWatcher{}
	gs := &oc.Sampling{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling", gs, queryPath, false, false)
		qv := (&oc.QualifiedSampling{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *SamplingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling) bool) *oc.SamplingWatcher {
	t.Helper()
	return watch_SamplingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *SamplingPath) Await(t testing.TB, timeout time.Duration, val *oc.Sampling) *oc.QualifiedSampling {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSampling) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling to the batch object.
func (n *SamplingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *SamplingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling {
	t.Helper()
	c := &oc.CollectionSampling{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_SamplingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling) bool) *oc.SamplingWatcher {
	t.Helper()
	w := &oc.SamplingWatcher{}
	structs := map[string]*oc.Sampling{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Sampling{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSampling{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *SamplingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling) bool) *oc.SamplingWatcher {
	t.Helper()
	return watch_SamplingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling to the batch object.
func (n *SamplingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_SflowPath) Lookup(t testing.TB) *oc.QualifiedSampling_Sflow {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSampling_Sflow{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_SflowPath) Get(t testing.TB) *oc.Sampling_Sflow {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_SflowPathAny) Lookup(t testing.TB) []*oc.QualifiedSampling_Sflow {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling_Sflow
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling_Sflow{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow with a ONCE subscription.
func (n *Sampling_SflowPathAny) Get(t testing.TB) []*oc.Sampling_Sflow {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling_Sflow
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_SflowPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling_Sflow {
	t.Helper()
	c := &oc.CollectionSampling_Sflow{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling_Sflow) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSampling_Sflow{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Sampling_Sflow)))
		return false
	})
	return c
}

func watch_Sampling_SflowPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling_Sflow) bool) *oc.Sampling_SflowWatcher {
	t.Helper()
	w := &oc.Sampling_SflowWatcher{}
	gs := &oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow", gs, queryPath, false, false)
		qv := (&oc.QualifiedSampling_Sflow{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling_Sflow)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_SflowPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling_Sflow) bool) *oc.Sampling_SflowWatcher {
	t.Helper()
	return watch_Sampling_SflowPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_SflowPath) Await(t testing.TB, timeout time.Duration, val *oc.Sampling_Sflow) *oc.QualifiedSampling_Sflow {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSampling_Sflow) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow to the batch object.
func (n *Sampling_SflowPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_SflowPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling_Sflow {
	t.Helper()
	c := &oc.CollectionSampling_Sflow{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling_Sflow) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_SflowPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling_Sflow) bool) *oc.Sampling_SflowWatcher {
	t.Helper()
	w := &oc.Sampling_SflowWatcher{}
	structs := map[string]*oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling_Sflow", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSampling_Sflow{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling_Sflow)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_SflowPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling_Sflow) bool) *oc.Sampling_SflowWatcher {
	t.Helper()
	return watch_Sampling_SflowPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow to the batch object.
func (n *Sampling_SflowPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_AgentIdIpv4Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_AgentIdIpv4Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_AgentIdIpv4Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_AgentIdIpv4Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a ONCE subscription.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_AgentIdIpv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_AgentIdIpv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_AgentIdIpv4Path(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_AgentIdIpv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_AgentIdIpv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_AgentIdIpv4Path) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 to the batch object.
func (n *Sampling_Sflow_AgentIdIpv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_AgentIdIpv4PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertSampling_Sflow_AgentIdIpv4Path(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_AgentIdIpv4PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/state/agent-id-ipv4 to the batch object.
func (n *Sampling_Sflow_AgentIdIpv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_AgentIdIpv4Path extracts the value of the leaf AgentIdIpv4 from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_AgentIdIpv4Path(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AgentIdIpv4
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
