package qos

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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Classifier_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Classifier", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_Classifier_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Classifier_TypePath) Get(t testing.TB) oc.E_Input_Classifier_Type {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Input_Classifier_Type
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Classifier", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Classifier_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a ONCE subscription.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Get(t testing.TB) []oc.E_Input_Classifier_Type {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Input_Classifier_Type
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Classifier_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Input_Classifier_Type {
	t.Helper()
	c := &oc.CollectionE_Input_Classifier_Type{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Input_Classifier_Type) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_Classifier_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Input_Classifier_Type) bool) *oc.E_Input_Classifier_TypeWatcher {
	t.Helper()
	w := &oc.E_Input_Classifier_TypeWatcher{}
	gs := &oc.Qos_Interface_Output_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Classifier", gs, queryPath, true, false)
		return convertQos_Interface_Output_Classifier_TypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Input_Classifier_Type)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Classifier_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_Classifier_Type) bool) *oc.E_Input_Classifier_TypeWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_Classifier_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_Classifier_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Input_Classifier_Type) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Input_Classifier_Type) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type to the batch object.
func (n *Qos_Interface_Output_Classifier_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Input_Classifier_Type {
	t.Helper()
	c := &oc.CollectionE_Input_Classifier_Type{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Input_Classifier_Type) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_Classifier_Type) bool) *oc.E_Input_Classifier_TypeWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_Classifier_TypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier/state/type to the batch object.
func (n *Qos_Interface_Output_Classifier_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_Classifier_TypePath extracts the value of the leaf Type from its parent oc.Qos_Interface_Output_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Input_Classifier_Type.
func convertQos_Interface_Output_Classifier_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Classifier) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	qv := &oc.QualifiedE_Input_Classifier_Type{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_InterfaceRef {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_InterfaceRefPath) Get(t testing.TB) *oc.Qos_Interface_Output_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a ONCE subscription.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output_InterfaceRef {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Output_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Output_InterfaceRef)))
		return false
	})
	return c
}

func watch_Qos_Interface_Output_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_InterfaceRef) bool) *oc.Qos_Interface_Output_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Output_InterfaceRefWatcher{}
	gs := &oc.Qos_Interface_Output_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Interface_Output_InterfaceRef{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Output_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_InterfaceRef) bool) *oc.Qos_Interface_Output_InterfaceRefWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Output_InterfaceRef) *oc.QualifiedQos_Interface_Output_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Output_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/interface-ref to the batch object.
func (n *Qos_Interface_Output_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output_InterfaceRef {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_InterfaceRef) bool) *oc.Qos_Interface_Output_InterfaceRefWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_InterfaceRefPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/interface-ref to the batch object.
func (n *Qos_Interface_Output_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_InterfaceRef", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a ONCE subscription.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Output_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", gs, queryPath, true, false)
		return convertQos_Interface_Output_InterfaceRef_InterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface to the batch object.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/interface to the batch object.
func (n *Qos_Interface_Output_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Qos_Interface_Output_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_InterfaceRef) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_InterfaceRef", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a ONCE subscription.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_Interface_Output_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_InterfaceRef", gs, queryPath, true, false)
		return convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface to the batch object.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/interface-ref/state/subinterface to the batch object.
func (n *Qos_Interface_Output_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Qos_Interface_Output_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Interface_Output_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_InterfaceRef) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_MulticastBufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output", gs, queryPath, true, false)
		return convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_MulticastBufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_MulticastBufferAllocationProfilePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/state/multicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Output_MulticastBufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_MulticastBufferAllocationProfilePath extracts the value of the leaf MulticastBufferAllocationProfile from its parent oc.Qos_Interface_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_MulticastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MulticastBufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_Queue {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_QueuePath) Get(t testing.TB) *oc.Qos_Interface_Output_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a ONCE subscription.
func (n *Qos_Interface_Output_QueuePathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_QueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output_Queue {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output_Queue) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Output_Queue{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Output_Queue)))
		return false
	})
	return c
}

func watch_Qos_Interface_Output_QueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Queue) bool) *oc.Qos_Interface_Output_QueueWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Output_QueueWatcher{}
	gs := &oc.Qos_Interface_Output_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Queue", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Interface_Output_Queue{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Output_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_QueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Queue) bool) *oc.Qos_Interface_Output_QueueWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_QueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_QueuePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Output_Queue) *oc.QualifiedQos_Interface_Output_Queue {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Output_Queue) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/queues/queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue to the batch object.
func (n *Qos_Interface_Output_QueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_QueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output_Queue {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output_Queue) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Queue) bool) *oc.Qos_Interface_Output_QueueWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_QueuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue to the batch object.
func (n *Qos_Interface_Output_QueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_Queue_AvgQueueLenPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_AvgQueueLenPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_Queue_AvgQueueLenPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Output_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Output_Queue_AvgQueueLenPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_AvgQueueLenPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len to the batch object.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_AvgQueueLenPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/avg-queue-len to the batch object.
func (n *Qos_Interface_Output_Queue_AvgQueueLenPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_Queue_AvgQueueLenPath extracts the value of the leaf AvgQueueLen from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Output_Queue_AvgQueueLenPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AvgQueueLen
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_DroppedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_Queue_DroppedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_DroppedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_DroppedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_DroppedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_DroppedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_DroppedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_Queue_DroppedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Output_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Output_Queue_DroppedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_DroppedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_DroppedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_Queue_DroppedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts to the batch object.
func (n *Qos_Interface_Output_Queue_DroppedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_DroppedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_DroppedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_DroppedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/dropped-pkts to the batch object.
func (n *Qos_Interface_Output_Queue_DroppedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_Queue_DroppedPktsPath extracts the value of the leaf DroppedPkts from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Output_Queue_DroppedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DroppedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_Queue_MaxQueueLenPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_MaxQueueLenPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_Queue_MaxQueueLenPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Output_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Output_Queue_MaxQueueLenPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_MaxQueueLenPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len to the batch object.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_MaxQueueLenPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/max-queue-len to the batch object.
func (n *Qos_Interface_Output_Queue_MaxQueueLenPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_Queue_MaxQueueLenPath extracts the value of the leaf MaxQueueLen from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Output_Queue_MaxQueueLenPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxQueueLen
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_Queue_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Output_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Output_Queue_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_Queue_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name to the batch object.
func (n *Qos_Interface_Output_Queue_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/name to the batch object.
func (n *Qos_Interface_Output_Queue_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_Queue_QueueManagementProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_Queue_QueueManagementProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a ONCE subscription.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_Queue_QueueManagementProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Output_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Output_Queue_QueueManagementProfilePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_QueueManagementProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile to the batch object.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_Queue_QueueManagementProfilePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/queues/queue/state/queue-management-profile to the batch object.
func (n *Qos_Interface_Output_Queue_QueueManagementProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_Queue_QueueManagementProfilePath extracts the value of the leaf QueueManagementProfile from its parent oc.Qos_Interface_Output_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_Queue_QueueManagementProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_Queue) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.QueueManagementProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
