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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_AgentIdIpv6Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_AgentIdIpv6Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_AgentIdIpv6Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertSampling_Sflow_AgentIdIpv6Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a ONCE subscription.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_AgentIdIpv6Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_AgentIdIpv6Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Sampling_Sflow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_AgentIdIpv6Path(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_AgentIdIpv6Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_AgentIdIpv6Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_AgentIdIpv6Path) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 to the batch object.
func (n *Sampling_Sflow_AgentIdIpv6Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_AgentIdIpv6PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertSampling_Sflow_AgentIdIpv6Path(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_AgentIdIpv6PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/state/agent-id-ipv6 to the batch object.
func (n *Sampling_Sflow_AgentIdIpv6PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_AgentIdIpv6Path extracts the value of the leaf AgentIdIpv6 from its parent oc.Sampling_Sflow
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_AgentIdIpv6Path(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AgentIdIpv6
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_CollectorPath) Lookup(t testing.TB) *oc.QualifiedSampling_Sflow_Collector {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSampling_Sflow_Collector{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_CollectorPath) Get(t testing.TB) *oc.Sampling_Sflow_Collector {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_CollectorPathAny) Lookup(t testing.TB) []*oc.QualifiedSampling_Sflow_Collector {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSampling_Sflow_Collector
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSampling_Sflow_Collector{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector with a ONCE subscription.
func (n *Sampling_Sflow_CollectorPathAny) Get(t testing.TB) []*oc.Sampling_Sflow_Collector {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Sampling_Sflow_Collector
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_CollectorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling_Sflow_Collector {
	t.Helper()
	c := &oc.CollectionSampling_Sflow_Collector{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling_Sflow_Collector) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSampling_Sflow_Collector{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Sampling_Sflow_Collector)))
		return false
	})
	return c
}

func watch_Sampling_Sflow_CollectorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Collector) bool) *oc.Sampling_Sflow_CollectorWatcher {
	t.Helper()
	w := &oc.Sampling_Sflow_CollectorWatcher{}
	gs := &oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow_Collector", gs, queryPath, false, false)
		qv := (&oc.QualifiedSampling_Sflow_Collector{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling_Sflow_Collector)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_CollectorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Collector) bool) *oc.Sampling_Sflow_CollectorWatcher {
	t.Helper()
	return watch_Sampling_Sflow_CollectorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/collectors/collector with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_CollectorPath) Await(t testing.TB, timeout time.Duration, val *oc.Sampling_Sflow_Collector) *oc.QualifiedSampling_Sflow_Collector {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSampling_Sflow_Collector) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/collectors/collector failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector to the batch object.
func (n *Sampling_Sflow_CollectorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_CollectorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSampling_Sflow_Collector {
	t.Helper()
	c := &oc.CollectionSampling_Sflow_Collector{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSampling_Sflow_Collector) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_CollectorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Collector) bool) *oc.Sampling_Sflow_CollectorWatcher {
	t.Helper()
	w := &oc.Sampling_Sflow_CollectorWatcher{}
	structs := map[string]*oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Sampling_Sflow_Collector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling_Sflow_Collector", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSampling_Sflow_Collector{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSampling_Sflow_Collector)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_CollectorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSampling_Sflow_Collector) bool) *oc.Sampling_Sflow_CollectorWatcher {
	t.Helper()
	return watch_Sampling_Sflow_CollectorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector to the batch object.
func (n *Sampling_Sflow_CollectorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_Collector_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a ONCE subscription.
func (n *Sampling_Sflow_Collector_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow_Collector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_Collector_AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_Collector_AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/collectors/collector/state/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/address to the batch object.
func (n *Sampling_Sflow_Collector_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Sampling_Sflow_Collector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Sampling_Sflow_Collector", structs[pre], queryPath, true, false)
			qv := convertSampling_Sflow_Collector_AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/address to the batch object.
func (n *Sampling_Sflow_Collector_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_Collector_AddressPath extracts the value of the leaf Address from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_Collector_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
