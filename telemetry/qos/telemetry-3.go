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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Classifier_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Classifier_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier", gs, queryPath, true, false)
		return convertQos_Interface_Input_Classifier_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Classifier_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name to the batch object.
func (n *Qos_Interface_Input_Classifier_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/name to the batch object.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Classifier_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Classifier_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_TermPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_Classifier_Term {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier_Term", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_TermPath) Get(t testing.TB) *oc.Qos_Interface_Input_Classifier_Term {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_TermPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_Classifier_Term {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_Classifier_Term
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_TermPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_Classifier_Term {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_Classifier_Term
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_TermPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_Classifier_Term {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_Classifier_Term{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_Classifier_Term) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_Classifier_Term{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_Classifier_Term)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Classifier_TermPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier_Term) bool) *oc.Qos_Interface_Input_Classifier_TermWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_Classifier_TermWatcher{}
	gs := &oc.Qos_Interface_Input_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Interface_Input_Classifier_Term{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_Classifier_Term)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_TermPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier_Term) bool) *oc.Qos_Interface_Input_Classifier_TermWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_TermPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Classifier_TermPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_Classifier_Term) *oc.QualifiedQos_Interface_Input_Classifier_Term {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_Classifier_Term) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term to the batch object.
func (n *Qos_Interface_Input_Classifier_TermPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_TermPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_Classifier_Term {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_Classifier_Term{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_Classifier_Term) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_TermPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier_Term) bool) *oc.Qos_Interface_Input_Classifier_TermWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_TermPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term to the batch object.
func (n *Qos_Interface_Input_Classifier_TermPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_Term_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier_Term", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Classifier_Term_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_Term_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_Term_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_Term_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_Term_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_Term_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Classifier_Term_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", gs, queryPath, true, false)
		return convertQos_Interface_Input_Classifier_Term_IdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_Term_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_Term_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Classifier_Term_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id to the batch object.
func (n *Qos_Interface_Input_Classifier_Term_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_Term_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_Term_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_Term_IdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/id to the batch object.
func (n *Qos_Interface_Input_Classifier_Term_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Classifier_Term_IdPath extracts the value of the leaf Id from its parent oc.Qos_Interface_Input_Classifier_Term
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Classifier_Term_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier_Term) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier_Term", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Classifier_Term_MatchedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_Term_MatchedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Classifier_Term_MatchedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", gs, queryPath, true, false)
		return convertQos_Interface_Input_Classifier_Term_MatchedOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_Term_MatchedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets to the batch object.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_Term_MatchedOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-octets to the batch object.
func (n *Qos_Interface_Input_Classifier_Term_MatchedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Classifier_Term_MatchedOctetsPath extracts the value of the leaf MatchedOctets from its parent oc.Qos_Interface_Input_Classifier_Term
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Classifier_Term_MatchedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier_Term) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier_Term", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Classifier_Term_MatchedPacketsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_Term_MatchedPacketsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Classifier_Term_MatchedPacketsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier_Term", gs, queryPath, true, false)
		return convertQos_Interface_Input_Classifier_Term_MatchedPacketsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_Term_MatchedPacketsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets to the batch object.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_Term_MatchedPacketsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/terms/term/state/matched-packets to the batch object.
func (n *Qos_Interface_Input_Classifier_Term_MatchedPacketsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Classifier_Term_MatchedPacketsPath extracts the value of the leaf MatchedPackets from its parent oc.Qos_Interface_Input_Classifier_Term
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Classifier_Term_MatchedPacketsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier_Term) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedPackets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Classifier_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_TypePath) Get(t testing.TB) oc.E_Input_Classifier_Type {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Input_Classifier_Type
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Get(t testing.TB) []oc.E_Input_Classifier_Type {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Input_Classifier_Type
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Input_Classifier_Type {
	t.Helper()
	c := &oc.CollectionE_Input_Classifier_Type{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Input_Classifier_Type) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Classifier_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Input_Classifier_Type) bool) *oc.E_Input_Classifier_TypeWatcher {
	t.Helper()
	w := &oc.E_Input_Classifier_TypeWatcher{}
	gs := &oc.Qos_Interface_Input_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier", gs, queryPath, true, false)
		return convertQos_Interface_Input_Classifier_TypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Input_Classifier_Type)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_Classifier_Type) bool) *oc.E_Input_Classifier_TypeWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Classifier_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Input_Classifier_Type) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Input_Classifier_Type) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type to the batch object.
func (n *Qos_Interface_Input_Classifier_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Input_Classifier_Type {
	t.Helper()
	c := &oc.CollectionE_Input_Classifier_Type{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Input_Classifier_Type) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_Classifier_Type) bool) *oc.E_Input_Classifier_TypeWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Classifier_TypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/state/type to the batch object.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Classifier_TypePath extracts the value of the leaf Type from its parent oc.Qos_Interface_Input_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Input_Classifier_Type.
func convertQos_Interface_Input_Classifier_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier) *oc.QualifiedE_Input_Classifier_Type {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_MulticastBufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input", gs, queryPath, true, false)
		return convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_MulticastBufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_MulticastBufferAllocationProfilePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/state/multicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_MulticastBufferAllocationProfilePath extracts the value of the leaf MulticastBufferAllocationProfile from its parent oc.Qos_Interface_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_Queue {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_QueuePath) Get(t testing.TB) *oc.Qos_Interface_Input_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription.
func (n *Qos_Interface_Input_QueuePathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_QueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_Queue {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_Queue) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_Queue{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_Queue)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_QueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Queue) bool) *oc.Qos_Interface_Input_QueueWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_QueueWatcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Interface_Input_Queue{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_QueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Queue) bool) *oc.Qos_Interface_Input_QueueWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_QueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_QueuePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_Queue) *oc.QualifiedQos_Interface_Input_Queue {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_Queue) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue to the batch object.
func (n *Qos_Interface_Input_QueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_QueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_Queue {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_Queue) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Queue) bool) *oc.Qos_Interface_Input_QueueWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_QueuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue to the batch object.
func (n *Qos_Interface_Input_QueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_AvgQueueLenPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_AvgQueueLenPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_AvgQueueLenPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_AvgQueueLenPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_AvgQueueLenPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len to the batch object.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_AvgQueueLenPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/avg-queue-len to the batch object.
func (n *Qos_Interface_Input_Queue_AvgQueueLenPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_AvgQueueLenPath extracts the value of the leaf AvgQueueLen from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Queue_AvgQueueLenPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_DroppedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_DroppedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_DroppedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_DroppedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_DroppedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_DroppedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_DroppedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_DroppedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_DroppedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_DroppedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_DroppedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_DroppedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts to the batch object.
func (n *Qos_Interface_Input_Queue_DroppedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_DroppedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_DroppedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_DroppedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/dropped-pkts to the batch object.
func (n *Qos_Interface_Input_Queue_DroppedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_DroppedPktsPath extracts the value of the leaf DroppedPkts from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Queue_DroppedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_MaxQueueLenPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_MaxQueueLenPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_MaxQueueLenPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_MaxQueueLenPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_MaxQueueLenPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len to the batch object.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_MaxQueueLenPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/max-queue-len to the batch object.
func (n *Qos_Interface_Input_Queue_MaxQueueLenPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_MaxQueueLenPath extracts the value of the leaf MaxQueueLen from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Queue_MaxQueueLenPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name to the batch object.
func (n *Qos_Interface_Input_Queue_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/name to the batch object.
func (n *Qos_Interface_Input_Queue_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_QueueManagementProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_QueueManagementProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_QueueManagementProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_QueueManagementProfilePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_QueueManagementProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile to the batch object.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_QueueManagementProfilePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/queue-management-profile to the batch object.
func (n *Qos_Interface_Input_Queue_QueueManagementProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_QueueManagementProfilePath extracts the value of the leaf QueueManagementProfile from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Queue_QueueManagementProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_TransmitOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_TransmitOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_TransmitOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_TransmitOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_TransmitOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets to the batch object.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_TransmitOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-octets to the batch object.
func (n *Qos_Interface_Input_Queue_TransmitOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_TransmitOctetsPath extracts the value of the leaf TransmitOctets from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Queue_TransmitOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TransmitOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_TransmitPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_Queue_TransmitPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_TransmitPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_TransmitPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_TransmitPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_TransmitPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_TransmitPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_Queue_TransmitPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Queue", gs, queryPath, true, false)
		return convertQos_Interface_Input_Queue_TransmitPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_TransmitPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_TransmitPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_Queue_TransmitPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts to the batch object.
func (n *Qos_Interface_Input_Queue_TransmitPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_Queue_TransmitPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_Queue_TransmitPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_Queue_TransmitPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/queues/queue/state/transmit-pkts to the batch object.
func (n *Qos_Interface_Input_Queue_TransmitPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_Queue_TransmitPktsPath extracts the value of the leaf TransmitPkts from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_Queue_TransmitPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TransmitPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Get(t testing.TB) *oc.Qos_Interface_Input_SchedulerPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_SchedulerPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_SchedulerPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_SchedulerPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_SchedulerPolicy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_SchedulerPolicy) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_SchedulerPolicy{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_SchedulerPolicy)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_SchedulerPolicy) bool) *oc.Qos_Interface_Input_SchedulerPolicyWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_SchedulerPolicyWatcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Interface_Input_SchedulerPolicy{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_SchedulerPolicy)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_SchedulerPolicy) bool) *oc.Qos_Interface_Input_SchedulerPolicyWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_SchedulerPolicy) *oc.QualifiedQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_SchedulerPolicy) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_SchedulerPolicy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_SchedulerPolicy) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_SchedulerPolicy) bool) *oc.Qos_Interface_Input_SchedulerPolicyWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicyPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy", gs, queryPath, true, false)
		return convertQos_Interface_Input_SchedulerPolicy_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/state/name to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_SchedulerPolicy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_SchedulerPolicy_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPath) Get(t testing.TB) *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_SchedulerPolicy_Scheduler{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_SchedulerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) bool) *oc.Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) bool) *oc.Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_SchedulerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_SchedulerPolicy_Scheduler{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) bool) *oc.Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_SchedulerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_SchedulerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-octets to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath extracts the value of the leaf ConformingOctets from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConformingOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/conforming-pkts to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath extracts the value of the leaf ConformingPkts from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_ConformingPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConformingPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-octets to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath extracts the value of the leaf ExceedingOctets from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ExceedingOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
