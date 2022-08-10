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

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_Collector_NetworkInstancePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertSampling_Sflow_Collector_NetworkInstancePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a ONCE subscription.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_NetworkInstancePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow_Collector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_Collector_NetworkInstancePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_NetworkInstancePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance to the batch object.
func (n *Sampling_Sflow_Collector_NetworkInstancePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_NetworkInstancePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertSampling_Sflow_Collector_NetworkInstancePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_NetworkInstancePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/network-instance to the batch object.
func (n *Sampling_Sflow_Collector_NetworkInstancePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_Collector_NetworkInstancePath extracts the value of the leaf NetworkInstance from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSampling_Sflow_Collector_NetworkInstancePath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NetworkInstance
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_PacketsSentPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_Collector_PacketsSentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_PacketsSentPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_PacketsSentPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_PacketsSentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a ONCE subscription.
func (n *Sampling_Sflow_Collector_PacketsSentPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_PacketsSentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_PacketsSentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow_Collector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_Collector_PacketsSentPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_PacketsSentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_PacketsSentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_Collector_PacketsSentPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent to the batch object.
func (n *Sampling_Sflow_Collector_PacketsSentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_PacketsSentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_PacketsSentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertSampling_Sflow_Collector_PacketsSentPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_PacketsSentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_PacketsSentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/packets-sent to the batch object.
func (n *Sampling_Sflow_Collector_PacketsSentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_Collector_PacketsSentPath extracts the value of the leaf PacketsSent from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSampling_Sflow_Collector_PacketsSentPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.PacketsSent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Sampling_Sflow_Collector_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Sampling_Sflow_Collector{}
	md, ok := oc.Lookup(t, n, "Sampling_Sflow_Collector", goStruct, true, false)
	if ok {
		return convertSampling_Sflow_Collector_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Sampling_Sflow_Collector_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Sampling_Sflow_Collector_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Sampling_Sflow_Collector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Sampling_Sflow_Collector", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSampling_Sflow_Collector_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a ONCE subscription.
func (n *Sampling_Sflow_Collector_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Sampling_Sflow_Collector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Sampling_Sflow_Collector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSampling_Sflow_Collector_PortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_PortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Sampling_Sflow_Collector_PortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-sampling/sampling/sflow/collectors/collector/state/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/port to the batch object.
func (n *Sampling_Sflow_Collector_PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Sampling_Sflow_Collector_PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Sampling_Sflow_Collector_PortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
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
			qv := convertSampling_Sflow_Collector_PortPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-sampling/sampling/sflow/collectors/collector/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Sampling_Sflow_Collector_PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Sampling_Sflow_Collector_PortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-sampling/sampling/sflow/collectors/collector/state/port to the batch object.
func (n *Sampling_Sflow_Collector_PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSampling_Sflow_Collector_PortPath extracts the value of the leaf Port from its parent oc.Sampling_Sflow_Collector
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSampling_Sflow_Collector_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.Sampling_Sflow_Collector) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Port
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
