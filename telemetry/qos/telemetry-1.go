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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, true, false)
	if ok {
		return convertQos_Classifier_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription.
func (n *Qos_Classifier_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier", gs, queryPath, true, false)
		return convertQos_Classifier_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/state/name to the batch object.
func (n *Qos_Classifier_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/state/name to the batch object.
func (n *Qos_Classifier_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_NamePath extracts the value of the leaf Name from its parent oc.Qos_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_TermPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_TermPath) Get(t testing.TB) *oc.Qos_Classifier_Term {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_TermPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
func (n *Qos_Classifier_TermPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_TermPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term)))
		return false
	})
	return c
}

func watch_Qos_Classifier_TermPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term) bool) *oc.Qos_Classifier_TermWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_TermWatcher{}
	gs := &oc.Qos_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_TermPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term) bool) *oc.Qos_Classifier_TermWatcher {
	t.Helper()
	return watch_Qos_Classifier_TermPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_TermPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term) *oc.QualifiedQos_Classifier_Term {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term to the batch object.
func (n *Qos_Classifier_TermPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_TermPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_TermPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term) bool) *oc.Qos_Classifier_TermWatcher {
	t.Helper()
	return watch_Qos_Classifier_TermPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term to the batch object.
func (n *Qos_Classifier_TermPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_ActionsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_ActionsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
func (n *Qos_Classifier_Term_ActionsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ActionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Actions)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_ActionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions) bool) *oc.Qos_Classifier_Term_ActionsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_ActionsWatcher{}
	gs := &oc.Qos_Classifier_Term_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Actions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ActionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions) bool) *oc.Qos_Classifier_Term_ActionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ActionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_ActionsPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Actions) *oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Actions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions to the batch object.
func (n *Qos_Classifier_Term_ActionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ActionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ActionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions) bool) *oc.Qos_Classifier_Term_ActionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ActionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions to the batch object.
func (n *Qos_Classifier_Term_ActionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Actions_Remark
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Actions_Remark
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions_Remark {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions_Remark{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Actions_Remark)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_RemarkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool) *oc.Qos_Classifier_Term_Actions_RemarkWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Actions_RemarkWatcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Actions_Remark)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool) *oc.Qos_Classifier_Term_Actions_RemarkWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_RemarkPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark to the batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions_Remark {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions_Remark{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool) *oc.Qos_Classifier_Term_Actions_RemarkWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_RemarkPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark to the batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_Remark_SetDot1PPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDot1PPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDot1PPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_Remark_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDot1P
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_Remark_SetDscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDscpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_Remark_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_Remark_SetMplsTcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetMplsTc
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_TargetGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_TargetGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group to the batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_TargetGroupPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group to the batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_TargetGroupPath extracts the value of the leaf TargetGroup from its parent oc.Qos_Classifier_Term_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Actions_TargetGroupPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TargetGroup
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_ConditionsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_ConditionsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_ConditionsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
func (n *Qos_Classifier_Term_ConditionsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ConditionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_ConditionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions) bool) *oc.Qos_Classifier_Term_ConditionsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_ConditionsWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ConditionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions) bool) *oc.Qos_Classifier_Term_ConditionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ConditionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_ConditionsPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions) *oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions to the batch object.
func (n *Qos_Classifier_Term_ConditionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ConditionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ConditionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions) bool) *oc.Qos_Classifier_Term_ConditionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ConditionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions to the batch object.
func (n *Qos_Classifier_Term_ConditionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions_Ipv4)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Ipv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool) *oc.Qos_Classifier_Term_Conditions_Ipv4Watcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Ipv4Watcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool) *oc.Qos_Classifier_Term_Conditions_Ipv4Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool) *oc.Qos_Classifier_Term_Conditions_Ipv4Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv4Path(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4/state/destination-address to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4_DestinationAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Qos_Classifier_Term_Conditions_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_Ipv4_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
