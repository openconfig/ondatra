package interfaces

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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref to the batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref to the batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6Path) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6PathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6 {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6) bool) *oc.Interface_Subinterface_Ipv6Watcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6) bool) *oc.Interface_Subinterface_Ipv6Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6Path) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6) *oc.QualifiedInterface_Subinterface_Ipv6 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 to the batch object.
func (n *Interface_Subinterface_Ipv6Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6 {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6) bool) *oc.Interface_Subinterface_Ipv6Watcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv6{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6) bool) *oc.Interface_Subinterface_Ipv6Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 to the batch object.
func (n *Interface_Subinterface_Ipv6PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Address {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_AddressPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Address {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Address)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool) *oc.Interface_Subinterface_Ipv6_AddressWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_AddressWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool) *oc.Interface_Subinterface_Ipv6_AddressWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_AddressPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedInterface_Subinterface_Ipv6_Address {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address to the batch object.
func (n *Interface_Subinterface_Ipv6_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Address {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool) *oc.Interface_Subinterface_Ipv6_AddressWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_AddressWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address) bool) *oc.Interface_Subinterface_Ipv6_AddressWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address to the batch object.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_IpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_IpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_IpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_IpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/ip to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_Subinterface_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_OriginPath) Get(t testing.TB) oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_IpAddressOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_IpAddressOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_OriginPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_IpAddressOrigin) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_OriginPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_OriginPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_OriginPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/origin to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_Subinterface_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_IpAddressOrigin.
func convertInterface_Subinterface_Ipv6_Address_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_IpAddressOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/prefix-length to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_Subinterface_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_StatusPath) Lookup(t testing.TB) *oc.QualifiedE_Address_Status {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_StatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_StatusPath) Get(t testing.TB) oc.E_Address_Status {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_StatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Address_Status {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Address_Status
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_StatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_StatusPathAny) Get(t testing.TB) []oc.E_Address_Status {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Address_Status
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_StatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Address_Status {
	t.Helper()
	c := &oc.CollectionE_Address_Status{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Address_Status) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_StatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	w := &oc.E_Address_StatusWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_StatusPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Address_Status)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_StatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_StatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_StatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_Address_Status) *oc.QualifiedE_Address_Status {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Address_Status) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_StatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_StatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Address_Status {
	t.Helper()
	c := &oc.CollectionE_Address_Status{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Address_Status) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_StatusPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	w := &oc.E_Address_StatusWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_StatusPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Address_Status)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_StatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_StatusPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/state/status to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_StatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_StatusPath extracts the value of the leaf Status from its parent oc.Interface_Subinterface_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Address_Status.
func convertInterface_Subinterface_Ipv6_Address_StatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedE_Address_Status {
	t.Helper()
	qv := &oc.QualifiedE_Address_Status{
		Metadata: md,
	}
	val := parent.Status
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAcceptMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath extracts the value of the leaf AcceptMode from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AcceptMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAdvertisementInterval())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath extracts the value of the leaf AdvertisementInterval from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AdvertisementInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath extracts the value of the leaf CurrentPriority from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CurrentPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriorityDecrement())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath extracts the value of the leaf PriorityDecrement from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PriorityDecrement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/track-interface to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath extracts the value of the leaf TrackInterface from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.TrackInterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPreemptDelay())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt-delay to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath extracts the value of the leaf PreemptDelay from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PreemptDelay
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPreempt())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/preempt to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath extracts the value of the leaf Preempt from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Preempt
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriority())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/priority to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath extracts the value of the leaf Priority from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-address to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath extracts the value of the leaf VirtualAddress from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.VirtualAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-link-local to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath extracts the value of the leaf VirtualLinkLocal from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.VirtualLinkLocal
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/state/virtual-router-id to the batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath extracts the value of the leaf VirtualRouterId from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.VirtualRouterId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Autoconf
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Autoconf
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Autoconf{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Autoconf)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_AutoconfPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool) *oc.Interface_Subinterface_Ipv6_AutoconfWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_AutoconfWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Autoconf)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool) *oc.Interface_Subinterface_Ipv6_AutoconfWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_AutoconfPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf to the batch object.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Autoconf{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_AutoconfPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool) *oc.Interface_Subinterface_Ipv6_AutoconfWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_AutoconfWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Autoconf{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Autoconf)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf) bool) *oc.Interface_Subinterface_Ipv6_AutoconfWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_AutoconfPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf to the batch object.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetCreateGlobalAddresses())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Autoconf{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-global-addresses to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath extracts the value of the leaf CreateGlobalAddresses from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.CreateGlobalAddresses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetCreateTemporaryAddresses())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Autoconf{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/create-temporary-addresses to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath extracts the value of the leaf CreateTemporaryAddresses from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.CreateTemporaryAddresses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetTemporaryPreferredLifetime())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Autoconf{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-preferred-lifetime to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath extracts the value of the leaf TemporaryPreferredLifetime from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.TemporaryPreferredLifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetTemporaryValidLifetime())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Autoconf{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Autoconf{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/state/temporary-valid-lifetime to the batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath extracts the value of the leaf TemporaryValidLifetime from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.TemporaryValidLifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_CountersPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Counters {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_CountersPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_CountersPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Counters {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Ipv6_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Ipv6_Counters)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool) *oc.Interface_Subinterface_Ipv6_CountersWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_CountersWatcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool) *oc.Interface_Subinterface_Ipv6_CountersWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedInterface_Subinterface_Ipv6_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters to the batch object.
func (n *Interface_Subinterface_Ipv6_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Ipv6_Counters {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Ipv6_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool) *oc.Interface_Subinterface_Ipv6_CountersWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Ipv6_CountersWatcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Ipv6_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Ipv6_Counters) bool) *oc.Interface_Subinterface_Ipv6_CountersWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters to the batch object.
func (n *Interface_Subinterface_Ipv6_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-discarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InDiscardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_InDiscardedPktsPath extracts the value of the leaf InDiscardedPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_InDiscardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InDiscardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_InErrorPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_InErrorPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InErrorPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_InErrorPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InErrorPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_InErrorPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-error-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InErrorPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_InErrorPktsPath extracts the value of the leaf InErrorPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_InErrorPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InErrorPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_InForwardedOctetsPath extracts the value of the leaf InForwardedOctets from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_InForwardedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InForwardedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-forwarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InForwardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_InForwardedPktsPath extracts the value of the leaf InForwardedPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_InForwardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InForwardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_InOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_InOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_InPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_InPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_InPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_InPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/in-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-discarded-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutDiscardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath extracts the value of the leaf OutDiscardedPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_OutDiscardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutDiscardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-error-pkts to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutErrorPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_OutErrorPktsPath extracts the value of the leaf OutErrorPkts from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_OutErrorPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutErrorPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Counters", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Interface_Subinterface_Ipv6_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Subinterface_Ipv6_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Subinterface_Ipv6_Counters", structs[pre], queryPath, true, false)
			qv := convertInterface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/state/counters/out-forwarded-octets to the batch object.
func (n *Interface_Subinterface_Ipv6_Counters_OutForwardedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath extracts the value of the leaf OutForwardedOctets from its parent oc.Interface_Subinterface_Ipv6_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_Subinterface_Ipv6_Counters_OutForwardedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutForwardedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
