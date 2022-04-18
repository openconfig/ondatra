package lacp

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

// Lookup fetches the value at /openconfig-lacp/lacp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LacpPath) Lookup(t testing.TB) *oc.QualifiedLacp {
	t.Helper()
	goStruct := &oc.Lacp{}
	md, ok := oc.Lookup(t, n, "Lacp", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LacpPath) Get(t testing.TB) *oc.Lacp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LacpPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp with a ONCE subscription.
func (n *LacpPathAny) Get(t testing.TB) []*oc.Lacp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LacpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp {
	t.Helper()
	c := &oc.CollectionLacp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp)))
		return false
	})
	return c
}

func watch_LacpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	w := &oc.LacpWatcher{}
	gs := &oc.Lacp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp", gs, queryPath, false, false)
		qv := (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LacpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	return watch_LacpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LacpPath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp) *oc.QualifiedLacp {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp to the batch object.
func (n *LacpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LacpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp {
	t.Helper()
	c := &oc.CollectionLacp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LacpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	w := &oc.LacpWatcher{}
	structs := map[string]*oc.Lacp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLacp{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LacpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	return watch_LacpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp to the batch object.
func (n *LacpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_InterfacePath) Lookup(t testing.TB) *oc.QualifiedLacp_Interface {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_InterfacePath) Get(t testing.TB) *oc.Lacp_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedLacp_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription.
func (n *Lacp_InterfacePathAny) Get(t testing.TB) []*oc.Lacp_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_Interface {
	t.Helper()
	c := &oc.CollectionLacp_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_Interface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp_Interface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp_Interface)))
		return false
	})
	return c
}

func watch_Lacp_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_Interface) bool) *oc.Lacp_InterfaceWatcher {
	t.Helper()
	w := &oc.Lacp_InterfaceWatcher{}
	gs := &oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface", gs, queryPath, false, false)
		qv := (&oc.QualifiedLacp_Interface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_Interface) bool) *oc.Lacp_InterfaceWatcher {
	t.Helper()
	return watch_Lacp_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp_Interface) *oc.QualifiedLacp_Interface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp_Interface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface to the batch object.
func (n *Lacp_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_Interface {
	t.Helper()
	c := &oc.CollectionLacp_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_Interface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_Interface) bool) *oc.Lacp_InterfaceWatcher {
	t.Helper()
	w := &oc.Lacp_InterfaceWatcher{}
	structs := map[string]*oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLacp_Interface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_Interface) bool) *oc.Lacp_InterfaceWatcher {
	t.Helper()
	return watch_Lacp_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface to the batch object.
func (n *Lacp_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_IntervalPath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, false)
	if ok {
		return convertLacp_Interface_IntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Lacp_LacpPeriodType{
		Metadata: md,
	}).SetVal(goStruct.GetInterval())
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_IntervalPath) Get(t testing.TB) oc.E_Lacp_LacpPeriodType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpPeriodType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a ONCE subscription.
func (n *Lacp_Interface_IntervalPathAny) Get(t testing.TB) []oc.E_Lacp_LacpPeriodType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpPeriodType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpPeriodType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpPeriodType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpPeriodType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpPeriodType) bool) *oc.E_Lacp_LacpPeriodTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpPeriodTypeWatcher{}
	gs := &oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpPeriodType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpPeriodType) bool) *oc.E_Lacp_LacpPeriodTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_IntervalPath) Await(t testing.TB, timeout time.Duration, val oc.E_Lacp_LacpPeriodType) *oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lacp_LacpPeriodType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/state/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/interval to the batch object.
func (n *Lacp_Interface_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpPeriodType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpPeriodType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpPeriodType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpPeriodType) bool) *oc.E_Lacp_LacpPeriodTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpPeriodTypeWatcher{}
	structs := map[string]*oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_IntervalPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpPeriodType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpPeriodType) bool) *oc.E_Lacp_LacpPeriodTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/interval to the batch object.
func (n *Lacp_Interface_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_IntervalPath extracts the value of the leaf Interval from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpPeriodType.
func convertLacp_Interface_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpPeriodType{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_LacpModePath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, false)
	if ok {
		return convertLacp_Interface_LacpModePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Lacp_LacpActivityType{
		Metadata: md,
	}).SetVal(goStruct.GetLacpMode())
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_LacpModePath) Get(t testing.TB) oc.E_Lacp_LacpActivityType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_LacpModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpActivityType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_LacpModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a ONCE subscription.
func (n *Lacp_Interface_LacpModePathAny) Get(t testing.TB) []oc.E_Lacp_LacpActivityType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpActivityType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_LacpModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpActivityType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpActivityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpActivityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_LacpModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpActivityTypeWatcher{}
	gs := &oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_Interface_LacpModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpActivityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_LacpModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_LacpModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_LacpModePath) Await(t testing.TB, timeout time.Duration, val oc.E_Lacp_LacpActivityType) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lacp_LacpActivityType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode to the batch object.
func (n *Lacp_Interface_LacpModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_LacpModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpActivityType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpActivityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpActivityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_LacpModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpActivityTypeWatcher{}
	structs := map[string]*oc.Lacp_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_Interface", structs[pre], queryPath, true, false)
			qv := convertLacp_Interface_LacpModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpActivityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_LacpModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_LacpModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode to the batch object.
func (n *Lacp_Interface_LacpModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_LacpModePath extracts the value of the leaf LacpMode from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpActivityType.
func convertLacp_Interface_LacpModePath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpActivityType{
		Metadata: md,
	}
	val := parent.LacpMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
